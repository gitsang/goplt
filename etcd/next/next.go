package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"v2/log"

	etcd3 "go.etcd.io/etcd/clientv3"
)

type Watcher interface {
	Next() ([]*etcd3.Event, error)
	Close()
}

type watcher struct {
	cli    *etcd3.Client
	key    string
	opts   []etcd3.OpOption
	ctx    context.Context
	cancel context.CancelFunc
	wch    etcd3.WatchChan
	err    error
}

func (w *watcher) firstNext() ([]*etcd3.Event, error) {
	resp, err := w.cli.Get(w.ctx, w.key, w.opts...)
	if w.err = err; err != nil {
		return nil, err
	}

	events := make([]*etcd3.Event, 0, len(resp.Kvs))
	for _, kv := range resp.Kvs {
		u := new(etcd3.Event)
		u.Type = etcd3.EventTypePut
		u.Kv = kv
		events = append(events, u)
	}
	log.Info("firstNext get resp", zap.Any("update-events", events))

	opts := []etcd3.OpOption{
		etcd3.WithRev(resp.Header.Revision + 1),
		etcd3.WithPrefix(),
		etcd3.WithPrevKV(),
	}
	w.wch = w.cli.Watch(w.ctx, w.key, opts...)
	return events, nil
}

func (w *watcher) Next() ([]*etcd3.Event, error) {
	if w.wch == nil {
		log.Warn("w.wch is nil, return firstNext")
		return w.firstNext()
	}
	if w.err != nil {
		return nil, w.err
	}

	wr, ok := <-w.wch
	if !ok {
		log.Error("watch closed")
		w.err = fmt.Errorf("watch closed")
		return nil, w.err
	}
	if w.err = wr.Err(); w.err != nil {
		log.Error(zap.Error(w.err))
		return nil, w.err
	}

	for _, e := range wr.Events {
		log.Info("wch recv", zap.Any("event", *e))
	}
	return wr.Events, nil
}

func (w *watcher) Close() {
	log.Warn("w.Close")
	if w == nil {
		return
	}
	w.cancel()
}

func NewWatcher(pctx context.Context, cli *etcd3.Client,
	key string, opts ...etcd3.OpOption) (Watcher, error) {
	ctx, cancel := context.WithCancel(pctx)
	return &watcher{
		cli:    cli,
		key:    key,
		opts:   opts,
		ctx:    ctx,
		cancel: cancel,
	}, nil
}
