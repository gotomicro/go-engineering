package main

import (
	"context"
	"errors"

	"github.com/gotomicro/ego"
	"github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/core/eerrors"
	"github.com/gotomicro/ego/core/elog"
	"go-engineering/helloworld"
)

func main() {
	if err := ego.New().Invoker(
		invokerGrpc,
		callGrpc,
	).Run(); err != nil {
		elog.Error("startup", elog.FieldErr(err))
	}
}

var grpcComp helloworld.GreeterClient

func invokerGrpc() error {
	grpcConn := egrpc.Load("grpc.test").Build()
	grpcComp = helloworld.NewGreeterClient(grpcConn.ClientConn)
	return nil
}

func callGrpc() error {
	_, err := grpcComp.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "error",
	})
	if err != nil {
		egoErr := eerrors.FromError(err)
		// egoErr.Is(helloworld.ResourceErrNotFound()) 一样的
		if errors.Is(egoErr, helloworld.ResourceErrNotFound()) {
			// 你的业务处理逻辑
			elog.Warn("404 not found", elog.Any("code", egoErr.GetCode()), elog.Any("message", egoErr.GetMessage()), elog.Any("metadata", egoErr.GetMetadata()))
			return nil
		} else {
			elog.Error(err.Error(), elog.Any("code", egoErr.GetCode()), elog.Any("message", egoErr.GetMessage()), elog.Any("metadata", egoErr.GetMetadata()))
			return err
		}
	}
	return nil
}
