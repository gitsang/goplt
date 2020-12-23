package main

import (
	"context"
	"time"
	"v2/log"
	"v2/utils"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	etcd3 "go.etcd.io/etcd/clientv3"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 2 * time.Second
	endpoints      = []string{"etcd1.l7i.top:2379"}
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

	k := "/sang/key"

	// Watcher
	//go func() {
	//	w := cli.Watch(context.Background(), "", etcd3.WithPrefix())
	//	for resp := range w {
	//		for _, e := range resp.Events {
	//			log.Info("watch", zap.Any("type", e.Type))
	//		}
	//	}
	//}()

	for {
		// PUT
		{
			v := "value-" + utils.RandomString(5)
			resp, err := cli.Put(context.Background(), k, v, etcd3.WithPrevKV())
			if err != nil {
				log.Error(zap.Error(err))
			} else {
				log.Info(zap.Any("resp", *resp))
			}
		}

		// GET
		{
			resp, err := cli.Get(context.Background(), k,
				etcd3.WithPrefix(), etcd3.WithSerializable())
			if err != nil {
				log.Error(zap.Error(err))
			} else {
				log.Info(zap.Any("Kvs", resp.Kvs))
			}
		}
		time.Sleep(5 * time.Second)
	}
}
