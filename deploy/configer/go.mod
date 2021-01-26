module configor

go 1.14

replace (
	//github.com/jinzhu/configor v1.2.1 => ../configor
)

require (
	gitcode.yealink.com/server/server_framework/go-utils/ylog v0.0.0-20201029085057-d73975b4a55a
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/configor v1.2.1
	github.com/pingcap/log v0.0.0-20201112100606-8f1e84a3abc8
	go.uber.org/zap v1.16.0
	golang.org/x/net v0.0.0-20190620200207-3b0461eec859
	google.golang.org/grpc v1.35.0
)
