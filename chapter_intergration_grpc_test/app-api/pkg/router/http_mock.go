package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gotomicro/ego/server/egin"
)

type HTTPMock struct {
	*egin.Component
}

func NewHTTPMock() *HTTPMock {
	return &HTTPMock{
		Server(),
	}
}

func (m *HTTPMock) MockPostByStruct(uri string, param interface{}) (*http.Response, error) {
	postByte, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	return m.MockPost(uri, postByte)
}

func (m *HTTPMock) MockPost(uri string, param []byte) (*http.Response, error) {
	// 构造post请求
	req := httptest.NewRequest("POST", uri, bytes.NewReader(param))
	req.Header.Set("Content-Type", "application/json")

	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应handler接口
	m.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	return result, nil
}

func (m *HTTPMock) MockGet(uri string) *http.Response {
	// 构造get请求
	req := httptest.NewRequest("GET", uri, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	m.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()
	return result
}
