package router

import (
	"context"
	"errors"

	"github.com/gotomicro/ego/server/egrpc"
	"go-engineering/proto/pb/resource/v1"
	"go-engineering/resource-svc/pkg/invoker"
	"go-engineering/resource-svc/pkg/model/mysql"
	"gorm.io/gorm"
)

func Server() *egrpc.Component {
	GRPCServer := egrpc.Load("server.grpc").Build()
	resourcev1.RegisterResourceServer(GRPCServer, &resourceGrpcServer{})
	return GRPCServer
}

type resourceGrpcServer struct {
	resourcev1.UnimplementedResourceServer
}

func (resourceGrpcServer) List(ctx context.Context, req *resourcev1.ListRequest) (resp *resourcev1.ListResponse, err error) {
	var list mysql.Resources
	err = invoker.Db.WithContext(ctx).Find(&list).Error
	if err != nil {
		return nil, resourcev1.ResourceErrListMysql().WithMessage(err.Error())
	}
	return &resourcev1.ListResponse{
		List: list.ToListPb(),
	}, nil
}

func (resourceGrpcServer) Detail(ctx context.Context, req *resourcev1.DetailRequest) (resp *resourcev1.DetailResponse, err error) {
	var info mysql.Resource
	err = invoker.Db.WithContext(ctx).Where("id = ?", req.Id).First(&info).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, resourcev1.ResourceErrInfoMysql().WithMessage(err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, resourcev1.ResourceErrNotFound().WithMessage(err.Error())
	}
	return info.ToDetailPb(), nil
}
