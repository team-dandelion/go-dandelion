package http

import (
	"fmt"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

type HttpServer struct {
	router *routing.Router
	port   int32
}

func New(port int32) *HttpServer {
	router := routing.New()
	router.Use(middlewareSentinel(), middlewareRequestLink(), middlewareCustomError())
	return &HttpServer{
		router: router,
		port:   port,
	}
}

func (hs *HttpServer) Router() *routing.RouteGroup {
	return hs.router.Group("/api")
}

func (hs *HttpServer) Port() int32 {
	return hs.port
}

func (hs *HttpServer) RegisterRoute(prefix string, routes []Route, middlewares ...routing.Handler) {
	router := hs.router.Group(prefix)
	router.Use(middlewares...)
	for _, route := range routes {
		switch route.Method {
		case GET:
			router.Get(route.Path, route.Handler)
		case POST:
			router.Post(route.Path, route.Handler)
		case PUT:
			router.Put(route.Path, route.Handler)
		case DELETE:
			router.Delete(route.Path, route.Handler)
		}
	}
}

func (hs *HttpServer) Server() {
	fasthttp.ListenAndServe(fmt.Sprintf(":%d", hs.port), hs.router.HandleRequest)
}
