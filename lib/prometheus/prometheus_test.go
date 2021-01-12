package prometheus

import "github.com/grpc-ecosystem/go-grpc-prometheus"
...
// Initialize your gRPC server's interceptor.
myServer := grpc.NewServer(
grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
)
// Register your gRPC service implementations.
myservice.RegisterMyServiceServer(s.server, &myServiceImpl{})
// After all your registrations, make sure all of the Prometheus metrics are initialized.
grpc_prometheus.Register(myServer)
// Register Prometheus metrics handler.
http.Handle("/metrics", promhttp.Handler())
...
