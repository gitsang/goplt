
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

## bbolt import err

```
go: go.etcd.io/etcd/clientv3/snapshot imports
	github.com/coreos/bbolt: github.com/coreos/bbolt@v1.3.5: parsing go.mod:
	module declares its path as: go.etcd.io/bbolt
	        but was required as: github.com/coreos/bbolt
```

### fix

```go.mod
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5
replace go.etcd.io/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.5
```
