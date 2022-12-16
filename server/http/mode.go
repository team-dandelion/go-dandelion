package http

import routing "github.com/qiangxue/fasthttp-routing"

type MethodType string

const (
	GET MethodType = "GET"
	POST MethodType = "POST"
	PUT MethodType = "PUT"
	DELETE MethodType = "DELETE"
)

type Route struct {
	Path string
	Method MethodType
	Handler func (ctx *routing.Context) error
}
