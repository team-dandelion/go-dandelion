package application

import (
	"github.com/team-dandelion/go-dandelion/config"
	"github.com/team-dandelion/go-dandelion/server/http"
)

var httpServer *http.HttpServer

func initHttpServer() {
	if config.Conf.HttpServer == nil {
		return
	}
	httpServer = http.New(config.Conf.HttpServer.Port)
}

func HttpServer() *http.HttpServer {
	if httpServer == nil {
		panic("未对httpServer进行配置")
	}
	return httpServer
}
