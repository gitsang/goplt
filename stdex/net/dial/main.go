package main

import (
	"context"
	"net"
	"time"

	log "github.com/gitsang/golog"
)

func main() {
	addr := "cn.ymw.pp.ua:9876"
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	var d net.Dialer
	_, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		log.Info("dial failed")
		return
	}
}
