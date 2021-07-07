package main

import (
	"encoding/json"
	"time"

	log "github.com/gitsang/golog"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
)

var origin = "http://10.200.112.67:1001/"
var url = "ws://10.200.112.67:1001/ws"

type Req struct {
	Method string `json:"method"`
	Msg    string `json:"msg"`
}

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Error("dial failed", zap.Error(err))
	}
	defer func() {
		_ = ws.Close()
	}()

	// make message
	req := Req{
		Method: "TIME",
		Msg:    "HELLO (GO)",
	}
	data, err := json.Marshal(req)
	if err != nil {
		log.Error("marshal failed", zap.Error(err))
		return
	}

	// send
	go func() {
		for {
			_, err = ws.Write(data)
			if err != nil {
				log.Error("write msg failed", zap.Error(err))
			}
			log.Info("send success", zap.ByteString("msg", data))
			time.Sleep(1 * time.Second)
		}
	}()

	// read
	go func() {
		for {
			var msgRead = make([]byte, 128)
			length, err := ws.Read(msgRead)
			msgRead = msgRead[0:length]
			if err != nil {
				log.Error("read msg failed", zap.Error(err))
			}
			log.Info("receive success", zap.ByteString("msg", msgRead), zap.Int("len", length))
		}
	}()

	select {}
}
