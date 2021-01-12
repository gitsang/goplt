package prometheus

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
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
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(),
			grpc_opentracing.StreamServerInterceptor(),
			grpc_prometheus.StreamServerInterceptor,
			//grpc_zap.StreamServerInterceptor(zapLogger),
			//grpc_auth.StreamServerInterceptor(myAuthFunction),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			//grpc_zap.UnaryServerInterceptor(zapLogger),
			//grpc_auth.UnaryServerInterceptor(myAuthFunction),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	//s := grpc.NewServer(
	//	grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
	//	grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	//)

	RegisterPingPongServer(s, &Server{})
	return s
}

func TestPrometheus(t *testing.T) {
	s := initGrpcServer()
	grpc_prometheus.Register(s)
	http.Handle("/metric", promhttp.Handler())
	_ = http.ListenAndServe(":8080", nil)
}
