package main

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type HelloReq struct {
	g.Meta `path:"/" method:"get" tags:"Test" summary:"Hello world test case"`
	Name   string `v:"required" dc:"姓名"` // 名称
	Age    int    `v:"required" dc:"年龄"` // 年龄
}

type HelloRes struct {
	Content string `json:"content" dc:"返回结果"`
}

type Hello struct{}

func (h Hello) Say(ctx context.Context, req *HelloReq) (res *HelloRes, err error) {
	res = &HelloRes{
		Content: fmt.Sprintf(
			"Hello, %s! Your Age is %d",
			req.Name, req.Age),
	}
	return
}

func ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.Write("error: ", err.Error())
		return
	}
}

func ResponseMiddleware(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		msg string
		res = r.GetHandlerResponse()
		err = r.GetError()
	)

	if err != nil {
		msg = err.Error()
	} else {
		msg = "ok"
	}
	r.Response.WriteJson(Response{
		Message: msg,
		Data:    res,
	})
}

type Response struct {
	Code    int         `json:"code" dc:"状态码"`
	Message string      `json:"message" dc:"信息提示"`
	Data    interface{} `json:"data" dc:"结果数据"`
}

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(ResponseMiddleware)
		group.Middleware(ErrorHandler)
		group.Bind(new(Hello))
	})
	s.SetOpenApiPath("/api.json")
	s.SetSwaggerPath("/swagger")
	s.SetPort(8000)
	s.Run()
}
