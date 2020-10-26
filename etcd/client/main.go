package main

import (
	"context"
	"fmt"
	"time"

	etcd3 "go.etcd.io/etcd/clientv3"
)

func main() {
	cli, err := etcd3.New(etcd3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/csh/key1", "value1")
	cancel()
	if err != nil {
		panic(err)
	}

	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "/csh/key1")
	cancel()
	if err != nil {
		panic(err)
	}
	for _, v := range resp.Kvs {
		fmt.Printf("%s:%s\n", v.Key, v.Value)
	}
}
