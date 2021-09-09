package router

import (
	"context"
	"log"
	"net"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	cegrpc "github.com/gotomicro/ego/client/egrpc"
	"github.com/gotomicro/ego/core/econf"
	"github.com/gotomicro/ego/core/eerrors"
	"github.com/gotomicro/ego/server/egrpc"
	"github.com/stretchr/testify/assert"
	resourcev1 "go-engineering/proto/pb/resource/v1"
	"go-engineering/resource-svc/pkg/invoker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

var svc *egrpc.Component

func init() {
	conf := `
[server.grpc]
network = "bufnet" # 使用bufnet模式的测试gRPC服务
[mysql.resource]
connMaxLifetime = "300s"
debug = true
dsn = "root:root@tcp(127.0.0.1:3306)/go-engineering?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=Local&timeout=1s&readTimeout=3s&writeTimeout=3s"
maxIdleConns = 50
maxOpenConns = 100
`
	// 加载配置
	err := econf.LoadFromReader(strings.NewReader(conf), toml.Unmarshal)
	if err != nil {
		log.Fatalf("init exited with error: %v", err)
	}

	// 初始化MySQL
	err = invoker.Init()
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

func bufDialer(context.Context, string) (net.Conn, error) {
	// 从测试gRPC服务获得listener
	return svc.Listener().(*bufconn.Listener).Dial()
}

func TestList(t *testing.T) {
	resourceClient := cegrpc.DefaultContainer().Build(cegrpc.WithDialOption(grpc.WithContextDialer(bufDialer)))
	ctx := context.Background()
	client := resourcev1.NewResourceClient(resourceClient.ClientConn)
	resp, err := client.List(ctx, &resourcev1.ListRequest{})
	assert.NoError(t, err)
	assert.Equal(t, int64(1), resp.List[0].Id)
	assert.Equal(t, "测试文章", resp.List[0].Title)
	assert.Equal(t, "ego", resp.List[0].Nickname)
	log.Printf("Response: %+v", resp)
}

func TestDetailOK(t *testing.T) {
	resourceClient := cegrpc.DefaultContainer().Build(cegrpc.WithDialOption(grpc.WithContextDialer(bufDialer)))
	ctx := context.Background()
	client := resourcev1.NewResourceClient(resourceClient.ClientConn)
	resp, err := client.Detail(ctx, &resourcev1.DetailRequest{
		Id: 1,
	})
	assert.NoError(t, err)
	assert.Equal(t, "测试文章", resp.Title)
	assert.Equal(t, "ego", resp.Nickname)
	assert.Equal(t, "测试文章内容", resp.Content)
	log.Printf("Response: %+v", resp)
}

func TestDetailNotFound(t *testing.T) {
	resourceClient := cegrpc.DefaultContainer().Build(cegrpc.WithDialOption(grpc.WithContextDialer(bufDialer)))
	ctx := context.Background()
	client := resourcev1.NewResourceClient(resourceClient.ClientConn)
	resp, err := client.Detail(ctx, &resourcev1.DetailRequest{
		Id: 2,
	})
	// 因为grpc的error是从远程调用得到的，必须通过errors.FromError，转为统一错误码，再来判定error的根因
	assert.ErrorIs(t, eerrors.FromError(err), resourcev1.ResourceErrNotFound())
	log.Printf("Response: %+v", resp)
}
