package main

import log "github.com/gitsang/golog"

func main() {
	log.Info("websocket server start")
	NewWebSocketServer("127.0.0.1:10081").Serve()
}
