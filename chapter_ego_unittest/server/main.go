package main

import (
	"egounittest/server/router"
	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/core/elog"
)

//  export EGO_DEBUG=true && go run main.go --config=config.toml
func main() {
	if err := ego.New().Serve(router.Server()).Run(); err != nil {
		elog.Panic("startup", elog.FieldErr(err))
	}
}
