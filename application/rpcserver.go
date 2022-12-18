package application

import (
	"context"
	"fmt"
	"github.com/gly-hub/go-dandelion/config"
	error_support "github.com/gly-hub/go-dandelion/error-support"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/server/http"
	"github.com/gly-hub/go-dandelion/server/rpcx"
	"github.com/gly-hub/go-dandelion/tools/ip"
	"github.com/gly-hub/go-dandelion/tools/stringx"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/smallnest/rpcx/share"
	"reflect"
)

var rpc *rpcServer

type rpcServer struct {
	ServerName string
	BasePath   string
	Etcd       []string
	Addr       string
	Port       int
	server     *rpcx.RPCxServer
	client     *rpcx.RPCXClient
}

func initRpcServer() {
	if config.Conf.RpcServer == nil && config.Conf.RpcServer.Etcd != nil {
		return
	}
	rpc = &rpcServer{
		ServerName: config.Conf.RpcServer.ServerName,
		BasePath:   config.Conf.RpcServer.BasePath,
		Etcd:       config.Conf.RpcServer.Etcd,
		Addr:       config.Conf.RpcServer.Addr,
		Port:       config.Conf.RpcServer.Port,
	}
	rpcx.SetBase(rpc.BasePath, rpc.Etcd)
	if config.Conf.RpcServer.Model != 0 {
		rpc.client = rpcx.NewRPCXClient(config.Conf.RpcServer.Model)
	}
}

// RpcCall rpc请求
func RpcCall(serverName, funcName string, args interface{}, reply interface{}, ctx ...context.Context)error{
	if rpc.client == nil {
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
		"request_id":  stringx.Strval(logger.GetRequestId()),
		"client_name": rpc.ServerName,
		"content":     content,
	}
	c = context.WithValue(c, share.ReqMetaDataKey, requestHeader)
	err := rpc.client.GetClient().Call(c, serverName, funcName, args, reply)
	if err != nil {
		return &error_support.Error{Code: 5001, Msg:  "服务器异常"}
	}

	rv := reflect.ValueOf(reply)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.FieldByName("Code").Int() != int64(0) {
		return &error_support.Error{Code: int(rv.FieldByName("Code").Int()), Msg:  rv.FieldByName("Msg").String()}
	}
	return nil
}

// SRpcCall rpc请求拓展
func SRpcCall(ctx *routing.Context, serverName, funcName string, args interface{}, reply interface{})error{
	if rpc.client == nil{
		panic("请配置rpcx参数")
	}
	var hc http.HttpController
	if err := hc.ReadJson(ctx, args); err != nil{
		return hc.Fail(ctx, &error_support.Error{Code: 5000, Msg:  "数据解析失败"})
	}

	content, _ := jsoniter.MarshalToString(args)
	requestHeader := map[string]string{
		"request_id":  stringx.Strval(logger.GetRequestId()),
		"client_name": rpc.ServerName,
		"content":     content,
	}
	c := context.Background()
	c = context.WithValue(c, share.ReqMetaDataKey, requestHeader)
	err := rpc.client.GetClient().Call(c, serverName, funcName, args, reply)
	if err != nil {
		return hc.Fail(ctx, &error_support.Error{Code: 5001, Msg:  "服务器异常"})
	}

	rv := reflect.ValueOf(reply)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.FieldByName("Code").Int() != int64(0) {
		return hc.Fail(ctx, &error_support.Error{Code: int(rv.FieldByName("Code").Int()), Msg:  rv.FieldByName("Msg").String()})
	}

	return hc.Success(ctx, reply, "")
}

func RpcServer(handler interface{}, auth ...rpcx.AuthCall) {
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

