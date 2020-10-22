
## grpc err

```
# go run main.go
# github.com/coreos/etcd/clientv3/balancer/resolver/endpoint
/root/go/pkg/mod/github.com/coreos/etcd@v3.3.25+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:114:78: undefined: resolver.BuildOption
/root/go/pkg/mod/github.com/coreos/etcd@v3.3.25+incompatible/clientv3/balancer/resolver/endpoint/endpoint.go:182:31: undefined: resolver.ResolveNowOption
# github.com/coreos/etcd/clientv3/balancer/picker
/root/go/pkg/mod/github.com/coreos/etcd@v3.3.25+incompatible/clientv3/balancer/picker/err.go:37:44: undefined: balancer.PickOptions
/root/go/pkg/mod/github.com/coreos/etcd@v3.3.25+incompatible/clientv3/balancer/picker/roundrobin_balanced.go:55:54: undefined: balancer.PickOptions
```

### fix

```bash
go mod edit -require=google.golang.org/grpc@v1.26.0
go get -u -x google.golang.org/grpc@v1.26.0
go run main.go
```
