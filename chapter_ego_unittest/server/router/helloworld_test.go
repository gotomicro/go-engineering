// Code generated by protoc-gen-go-test. DO NOT EDIT.

package router

import (
	"context"
	"errors"
	"testing"

	"egounittest/protos/helloworld"
	cegrpc "github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/core/eerrors"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the ego package it is being compiled against.

// TestSayHello generated by protoc-gen-go-test, you can fill test logic by yourself.
// @Override=true
func TestSayHello(t *testing.T) {
	cli := helloworld.NewGreeterClient(cegrpc.DefaultContainer().Build(cegrpc.WithBufnetServerListener(svc.Listener())).ClientConn)
	ctx := context.Background()
	tests := []struct {
		name    string
		req     *helloworld.HelloRequest
		wantRes *helloworld.HelloResponse
		wantErr *eerrors.EgoError
	}{
		// TODO: Add or modify test cases.
		{"正常情况", &helloworld.HelloRequest{}, &helloworld.HelloResponse{
			Message: "Hello EGO",
		}, nil},
		{"错误情况", &helloworld.HelloRequest{
			Name: "error",
		}, nil, helloworld.ResourceErrNotFound().(*eerrors.EgoError)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := cli.SayHello(ctx, tt.req)
			assert.True(t, errors.Is(eerrors.FromError(err), tt.wantErr))
			assert.True(t, proto.Equal(tt.wantRes, res))
			t.Logf("res: %+v", res)
		})
	}
}
