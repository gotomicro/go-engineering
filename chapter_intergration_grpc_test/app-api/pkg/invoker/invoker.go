package invoker

import (
	"github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/core/elog"
	resourcev1 "go-engineering/proto/pb/resource/v1"
)

var (
	Logger       = elog.DefaultLogger
	ResourceGrpc resourcev1.ResourceClient
)

func Init() error {
	// 如果使用k8s协议，那么必须用k8s的配置
	//registry.DefaultContainer().Build(registry.WithClient(ek8s.Load("k8s").Build()))
	Logger = elog.DefaultLogger
	ResourceGrpc = resourcev1.NewResourceClient(egrpc.Load("grpc.resource").Build().ClientConn)
	return nil
}
