package code

import (
	"fmt"
)

func ServiceApi() string {
	return `package service

type RpcApi struct {

}`
}

func Route() string {
	return fmt.Sprintf(`package route

import (
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gly-hub/go-dandelion/config"
	"github.com/gly-hub/go-dandelion/server/http"
	routingSwagger "github.com/gly-hub/go-dandelion/swagger"
)

func InitRoute() {
	baseRouter := application.HttpServer().Router()
	if config.GetEnv() != "production" {
		// 注册swagger
		baseRouter.Get("/swagger/*", routingSwagger.WrapHandler)
		http.LogIgnoreResult(".*?/swagger/.*?") // 忽略swagger响应值
	}

	// 可在该处注册相关子集路由 TODO
}
`)
}
