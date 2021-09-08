package invoker

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/server/egrpc"
)

var (
	Logger     = elog.DefaultLogger
	Db         *egorm.Component
	GRPCServer *egrpc.Component
)

func Init() error {
	Logger = elog.DefaultLogger
	Db = egorm.Load("mysql.resource").Build()
	GRPCServer = egrpc.Load("server.grpc").Build()
	return nil
}
