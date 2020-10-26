module v2

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5

replace go.etcd.io/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.5

require (
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/prometheus/client_golang v1.8.0 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	google.golang.org/grpc v1.33.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)
