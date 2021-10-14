// Code generated by protoc-gen-go-test. DO NOT EDIT.

package router

import (
	"context"
	"log"
	"net"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/server/egrpc"
	"google.golang.org/grpc/test/bufconn"
)

// svc generated by protoc-gen-go-test, you should not edit it.
// @Overrid=true
var svc *egrpc.Component

// init generated by protoc-gen-go-test, you can fill initial logic by yourself.
// @Override=true
func init() {
	conf := `
mode = "unittest"
[server.grpc]
network = "bufnet" # 使用bufnet模式的测试gRPC服务
enableAccessInterceptor = true
enableAccessInterceptorRes = true
enableAccessInterceptorReq = true
`
	// 加载配置
	err := econf.LoadFromReader(strings.NewReader(conf), toml.Unmarshal)
	if err != nil {
		log.Fatalf("init exited with error: %v", err)
	}

	// 初始化bufnet gRPC的测试服务
	svc = Server()

	err = svc.Init()
	if err != nil {
		log.Fatalf("init server with error: %v", err)
	}

	go func() {
		// 启动服务
		if err = svc.Start(); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

// bufDialer generated by protoc-gen-go-test, you should not edit it.
// @Override=true
func bufDialer(context.Context, string) (net.Conn, error) {
	return svc.Listener().(*bufconn.Listener).Dial()
}
