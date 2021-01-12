package main

import (
	"context"
	etcd3 "go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
	"v2/log"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"etcd1.l7i.top:2379"}
	key            = "/sang/key"
)

func main() {
	log.InitLogger(zapcore.DebugLevel, "watcher.debug.log")
	defer log.Sync()

	cli, err := etcd3.New(etcd3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Error(zap.Error(err))
	}
	defer cli.Close()

	for {
		w, err := NewWatcher(context.Background(), cli, key,
			etcd3.WithPrefix(), etcd3.WithSerializable())
		if err != nil {
			log.Error("new watcher failed", zap.Error(err))
			continue
		}

		for {
			events, err := w.Next()
			if err != nil {
				log.Error("watch next failed", zap.Error(err))
				break
			}
			for _, e := range events {
				log.Info(zap.Any("event", *e))
			}
		}
		time.Sleep(time.Second)
	}
}
