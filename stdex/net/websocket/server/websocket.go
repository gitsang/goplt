package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	log "github.com/gitsang/golog"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WebSocketServer struct {
	upgrader *websocket.Upgrader
	addr     string
}

func NewWebSocketServer(addr string) *WebSocketServer {
	return &WebSocketServer{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  4096,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				if r.Method != "GET" {
					log.Error("method is not GET")
					return false
				}
				if r.URL.Path != "/ws" {
					log.Error("path error", zap.String("path", r.URL.Path))
					return false
				}
				return true
			},
			EnableCompression: false,
		},
		addr: addr,
	}
}

func (ws *WebSocketServer) send(ctx context.Context, conn *websocket.Conn) {
	for {
		select {
		case <-ctx.Done():
			log.Warn("connect closed")
			return
		case <-time.After(time.Second):
			log.Info("sending....")
			data := fmt.Sprintf("websocket server hello %v", time.Now().UnixNano())
			err := conn.WriteMessage(1, []byte(data))
			if err != nil {
				fmt.Println("send msg failed", err)
				return
			}
		}
	}
}

func (ws *WebSocketServer) connHandle(conn *websocket.Conn) {
	defer func() {
		_ = conn.Close()
	}()

	ctx, cancel := context.WithCancel(context.Background())
	go ws.send(ctx, conn)

	for {
		_ = conn.SetReadDeadline(time.Now().Add(time.Millisecond * time.Duration(5000)))
		_, msg, err := conn.ReadMessage()
		if err != nil {
			cancel()

			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					log.Error("read message timeout", zap.String("remote", conn.RemoteAddr().String()))
					return
				}
			}

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Error("read message failed",
					zap.String("remote", conn.RemoteAddr().String()),
					zap.Error(err))
			}

			return
		}

		log.Info("receive msg", zap.ByteString("msg", msg))
	}
}

func (ws *WebSocketServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("websocket upgrade failed", zap.Error(err))
		return
	}
	log.Info("websocket client connect",
		zap.String("network", conn.RemoteAddr().Network()),
		zap.String("addr", conn.RemoteAddr().String()))

	respStr := "Hello WebSocket init: " + ws.addr
	l, err := w.Write([]byte(respStr))
	if err != nil {
		log.Error("write response failed", zap.Error(err))
	}
	log.Info("write response success", zap.Int("len", l))

	go ws.connHandle(conn)
}

func (ws *WebSocketServer) Serve() {
	_ = http.ListenAndServe(ws.addr, ws)
}

