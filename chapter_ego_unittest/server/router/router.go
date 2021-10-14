package router

import (
	"context"
	"fmt"

	"egounittest/protos/helloworld"
	"github.com/gotomicro/ego/server/egrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Server() *egrpc.Component {
	grpcCmp := egrpc.Load("server.grpc").Build()
	helloworld.RegisterGreeterServer(grpcCmp.Server, &Greeter{})
	return grpcCmp
}

// Greeter ...
type Greeter struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello ...
func (g Greeter) SayHello(context context.Context, request *helloworld.HelloRequest) (*helloworld.HelloResponse, error) {
	if request.Name == "error" {
		return nil, helloworld.ResourceErrNotFound().WithMessage("sayhello err").WithMetadata(map[string]string{
			"ego1": "haha",
			"ego2": "wawa",
		})
	}
	header := metadata.Pairs("x-header-key", "val")
	err := grpc.SendHeader(context, header)
	if err != nil {
		return nil, fmt.Errorf("set header fail, %w", err)
	}
	return &helloworld.HelloResponse{
		Message: "Hello EGO",
	}, nil
}
