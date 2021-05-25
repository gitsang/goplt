package main

import (
	"context"
	"net"
	"time"

	log "github.com/gitsang/golog"
	"go.uber.org/zap"
)

func main() {
	var err error

	_, err = net.Dial("tcp", "cn.ymw.pp.ua:9876")
	if err != nil {
		log.Error("dial failed", zap.Error(err))
	} else {
		log.Info("dial success")
	}

	var d net.Dialer
	_, err = d.Dial("tcp", "cn.ymw.pp.ua:9876")
	if err != nil {
		log.Error("dial failed", zap.Error(err))
	} else {
		log.Info("dial success")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err = d.DialContext(ctx, "tcp", "cn.ymw.pp.ua:9876")
	if err != nil {
		log.Error("dial context failed", zap.Error(err))
	} else {
		log.Info("dail context success")
	}
}
