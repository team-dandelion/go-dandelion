package application

import (
	"fmt"
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/toolbox/ip"
	"github.com/smallnest/rpcx/server"
	"github.com/team-dandelion/go-dandelion/config"
	"github.com/team-dandelion/go-dandelion/logger"
	"github.com/team-dandelion/go-dandelion/server/rpcx"
)

var (
	rpcServer  *rpcx.Server
	rpcClient  *RpcClient
	rpcPlugins []server.Plugin
)

type RpcClient struct {
	ClientName string
	ClientPool *rpcx.ClientPool
	HeaderFunc func(ctx *routing.Context, header map[string]string) map[string]string
}

func initRpcClient() {
	if config.Conf.RpcClient == nil {
		return
	}
	client, err := rpcx.NewRPCClient(rpcx.ClientConfig{
		ClientName:      config.Conf.RpcClient.ClientName,
		BasePath:        config.Conf.RpcClient.BasePath,
		RegisterPlugin:  rpcx.RegisterPluginType(config.Conf.RpcClient.RegisterPlugin),
		RegisterServers: config.Conf.RpcClient.RegisterServers,
		FailRetryModel:  rpcx.FailRetryModel(config.Conf.RpcClient.FailRetryModel),
		BalanceModel:    rpcx.BalanceModel(config.Conf.RpcClient.BalanceModel),
		PoolSize:        config.Conf.RpcClient.PoolSize,
	})
	if err != nil {
		panic(err)
	}
	rpcClient = &RpcClient{
		ClientName: config.Conf.RpcClient.ClientName,
		ClientPool: client,
	}
}

func RegisterHeaderFunc(f func(ctx *routing.Context, header map[string]string) map[string]string) {
	rpcClient.HeaderFunc = f
}

func RegisterRpcPlugin(plugins ...server.Plugin) {
	rpcPlugins = append(rpcPlugins, plugins...)
}

func RpcServer(handler interface{}, auth ...rpcx.AuthFunc) {
	if config.Conf.RpcServer == nil || config.Conf.RpcServer.ServerName == "" ||
		config.Conf.RpcServer.Port == 0 || config.Conf.RpcServer.RegisterServers == nil {
		logger.Error("请检查rpc服务配置")
		panic("请检查rpc服务配置")
	}

	var addr string
	if config.Conf.RpcServer.Addr != "" {
		addr = config.Conf.RpcServer.Addr
	} else {
		addr = ip.GetLocalHost()
	}
	var err error
	rpcServer, err = rpcx.NewRPCServer(rpcx.ServerConfig{
		ServerName:      config.Conf.RpcServer.ServerName,
		Addr:            fmt.Sprintf("%s:%d", addr, config.Conf.RpcServer.Port),
		BasePath:        config.Conf.RpcServer.BasePath,
		RegisterPlugin:  rpcx.RegisterPluginType(config.Conf.RpcServer.RegisterPlugin),
		RegisterServers: config.Conf.RpcServer.RegisterServers,
		Handle:          handler,
	}, rpcPlugins...)
	if err != nil {
		panic(err)
	}
	if len(auth) > 0 {
		rpcServer.RegisterAuthFunc(auth[0])
	}
	rpcServer.Start()
}

func GetRpcClient() *RpcClient {
	return rpcClient
}
