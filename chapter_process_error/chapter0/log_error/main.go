package main

import (
	"github.com/gotomicro/ego/core/elog"
)

func main() {
	elog.Error("grpc error", elog.FieldErrAny("param error"))
	select {}
}
