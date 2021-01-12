package prometheus

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"net/http"
	"testing"
)

type Server struct {
}

func (s *Server) Ping(ctx context.Context, pingMsg *PingMsg) (*PongMsg, error) {
	fmt.Println(pingMsg.Message)
	return &PongMsg{Message: "Pong"}, nil
}

func initGrpcServer() *grpc.Server {
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	RegisterPingPongServer(s, &Server{})
	return s
}

func TestPrometheus(t *testing.T) {
	s := initGrpcServer()
	grpc_prometheus.Register(s)
	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe(":8080", nil)
}
