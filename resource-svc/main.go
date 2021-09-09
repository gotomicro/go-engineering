package main

import (
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
	"github.com/gotomicro/ego/task/ejob"
	"go-engineering/resource-svc/pkg/invoker"
	"go-engineering/resource-svc/pkg/job"
	"go-engineering/resource-svc/pkg/router"
)

func main() {
	err := ego.New().
		Invoker(
			invoker.Init,
		).
		Job(
			ejob.Job("install", job.InstallComponent),
			ejob.Job("initialize", job.InitializeComponent),
		).
		Serve(
			router.Server(),
		).Run()
	if err != nil {
		elog.Panic("start up error: " + err.Error())
	}
}
