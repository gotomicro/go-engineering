package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"go-engineering/app-api/pkg/invoker"
	"go-engineering/app-api/pkg/router"
)

func main() {
	err := ego.New().
		Invoker(
			invoker.Init,
		).
		Serve(router.Server()).Run()
	if err != nil {
		elog.Panic("start up error: " + err.Error())
	}
}
