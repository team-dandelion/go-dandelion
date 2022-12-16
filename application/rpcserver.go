package application

import (
	"context"
	"fmt"
	"github.com/gly-hub/go-dandelion/config"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/server/rpcx"
	"github.com/gly-hub/go-dandelion/tools/ip"
	"github.com/gly-hub/go-dandelion/tools/stringx"
	jsoniter "github.com/json-iterator/go"
	"github.com/smallnest/rpcx/share"
)

var rpc *rpcServer

type rpcServer struct{
	ServerName string
	BasePath string
	Etcd []string
	Addr string
	Port int
	server *rpcx.RPCxServer
	client *rpcx.RPCXClient
}

func initRpcServer(){
	if config.Conf.RpcServer == nil && config.Conf.RpcServer.Etcd != nil {
		return
	}
	rpc = &rpcServer{
		ServerName: config.Conf.RpcServer.ServerName,
		BasePath: config.Conf.RpcServer.BasePath,
		Etcd: config.Conf.RpcServer.Etcd,
		Addr:       config.Conf.RpcServer.Addr,
		Port:       config.Conf.RpcServer.Port,
	}
	rpcx.SetBase(rpc.BasePath, rpc.Etcd)
	if config.Conf.RpcServer.Model != 0 {
		rpc.client = rpcx.NewRPCXClient(config.Conf.RpcServer.Model)
	}
}

func RpcCall(serverName, funcName string, args interface{}, reply interface{}, ctx ...context.Context)error{
	if rpc.client == nil{
		panic("请配置rpcx参数")
	}
	var c context.Context
	if len(ctx) > 0 {
		c = ctx[0]
	} else {
		c = context.Background()
	}

	content, _ := jsoniter.MarshalToString(args)
	requestHeader := map[string]string{
		"request_id": stringx.Strval(logger.GetRequestId()),
		"client_name": rpc.ServerName,
		"content": content,
	}
	c = context.WithValue(c, share.ReqMetaDataKey, requestHeader)
	return rpc.client.GetClient().Call(c, serverName, funcName, args, reply)
}

func RpcServer(handler interface{}, auth ...rpcx.AuthCall){
	if rpc == nil {
		panic("请配置rpcx参数")
	}

	var addr string
	if rpc.Addr != "" {
		addr = rpc.Addr
	} else {
		addr = ip.GetLocalHost()
	}
	rpc.server = rpcx.NewRpcxServer(rpc.ServerName, fmt.Sprintf("%s:%d", addr, rpc.Port), handler)
	if len(auth) > 0 {
		rpc.server.SetAuthCall(auth[0])
	}
	rpc.server.Start()
}

