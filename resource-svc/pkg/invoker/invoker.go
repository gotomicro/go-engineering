package invoker

import (
	"github.com/gotomicro/ego-component/egorm"
	"github.com/gotomicro/ego/core/elog"
)

var (
	Logger = elog.DefaultLogger
	Db     *egorm.Component
)

func Init() error {
	Logger = elog.DefaultLogger
	Db = egorm.Load("mysql.resource").Build()
	return nil
}
