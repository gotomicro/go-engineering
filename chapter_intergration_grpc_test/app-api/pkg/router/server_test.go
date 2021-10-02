package router

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/gotomicro/ego/core/econf"
	"github.com/stretchr/testify/assert"
	"go-engineering/app-api/pkg/invoker"
)

var svc *HTTPMock

func init() {
	conf := `
[server.http]
[grpc.resource]
addr = "127.0.0.1:9001"
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
	svc = NewHTTPMock()
}

func TestList(t *testing.T) {
	response := svc.MockGet("/list")
	// 读取响应body
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, `{"code":0,"data":{"list":[{"id":1,"title":"测试文章","nickname":"ego"}]}}`, string(body))
}

func TestDetailOk(t *testing.T) {
	response := svc.MockGet("/detail/1")
	// 读取响应body
	body, err := ioutil.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, `{"code":0,"data":{"title":"测试文章","nickname":"ego","content":"测试文章内容"}}`, string(body))
}

func TestDetailNotFound(t *testing.T) {
	response := svc.MockGet("/detail/2")
	// 读取响应body
	body, err := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.NoError(t, err)
	assert.Equal(t, `{"code":1,"msg":"获取详情数据失败"}`, string(body))
}
