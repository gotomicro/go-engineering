package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gotomicro/ego/core/eerrors"
	"github.com/gotomicro/ego/core/transport"
	"github.com/gotomicro/ego/server/egin"
	"go-engineering/app-api/pkg/invoker"
	resourcev1 "go-engineering/proto/pb/resource/v1"
)

func Server() *egin.Component {
	router := invoker.Gin
	router.Use(MockLogin())
	router.GET("/", helloEgo)
	router.GET("/list", resourceList)
	router.GET("/detail/:id", resourceDetail)
	return router
}

// MockLogin 模拟用户登陆，并传递uid的链路信息
func MockLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置链路uid，在远程调用中传递uid信息
		parentContext := transport.WithValue(c.Request.Context(), "X-Ego-Uid", 9527)
		c.Request = c.Request.WithContext(parentContext)
		fmt.Printf("	c.Request --------------->"+"%+v\n", transport.Value(c.Request.Context(), "X-Ego-Uid"))
		c.Set("X-Ego-Uid", 9527)
		c.Next()
	}
}

func helloEgo(ctx *gin.Context) {
	fmt.Printf("transport.CustomContextKeys()--------------->"+"%+v\n", transport.CustomContextKeys())
	ctx.String(http.StatusOK, "hello ego")
}

func resourceList(ctx *gin.Context) {
	list, err := invoker.ResourceGrpc.List(ctx.Request.Context(), &resourcev1.ListRequest{})
	if err != nil {
		nerr := eerrors.FromError(err)
		ctx.JSON(nerr.ToHTTPStatusCode(), gin.H{
			"msg": "获取列表数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": list,
	})

}

func resourceDetail(ctx *gin.Context) {
	list, err := invoker.ResourceGrpc.Detail(ctx.Request.Context(), &resourcev1.DetailRequest{})
	if err != nil {
		nerr := eerrors.FromError(err)
		ctx.JSON(nerr.ToHTTPStatusCode(), gin.H{
			"msg": "获取详情数据失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data": list,
	})
}
