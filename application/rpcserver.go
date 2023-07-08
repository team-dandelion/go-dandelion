package application

import (
	"context"
	"fmt"
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-dandelion/config"
	error_support "github.com/gly-hub/go-dandelion/error-support"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/server/http"
	"github.com/gly-hub/go-dandelion/server/rpcx"
	"github.com/gly-hub/go-dandelion/telemetry"
	"github.com/gly-hub/toolbox/ip"
	"github.com/gly-hub/toolbox/stringx"
	jsoniter "github.com/json-iterator/go"
	"github.com/smallnest/rpcx/server"
	_ "net/http/pprof"
	"reflect"
)

var (
	rpcServer  *rpcx.Server
	rpcClient  *RpcClient
	rpcPlugins []server.Plugin
)

type RpcClient struct {
	ClientName string
	clientPool *rpcx.ClientPool
	headerFunc func(ctx *routing.Context, header map[string]string) map[string]string
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
		clientPool: client,
	}
}

func RegisterHeaderFunc(f func(ctx *routing.Context, header map[string]string) map[string]string) {
	rpcClient.headerFunc = f
}

// RpcCall rpc请求
func RpcCall(ctx *routing.Context, serverName, funcName string, args interface{}, reply interface{}) error {
	if rpcClient.clientPool == nil {
		panic("请配置rpcx参数")
	}
	content, _ := jsoniter.MarshalToString(args)
	var traceId string
	if telemetry.GetSpanTraceId() != nil {
		traceId = telemetry.GetSpanTraceId().(string)
	}
	requestHeader := map[string]string{
		"request_id":    stringx.Strval(logger.GetRequestId()),
		"span_trace_id": traceId,
		"client_name":   rpcClient.ClientName,
		"content":       content,
	}

	requestHeader = rpcClient.headerFunc(ctx, requestHeader)
	c := rpcx.Header().Set(context.Background(), requestHeader)
	err := rpcClient.clientPool.Client().Call(c, serverName, funcName, args, reply)
	if err != nil {
		logger.Error("ServerName: ", serverName, ", FuncName: ", funcName, ", Err: ", err)
		return &error_support.Error{Code: 5001, Msg: "服务器异常"}
	}

	rv := reflect.ValueOf(reply)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.FieldByName("Code").Int() != int64(0) {
		return &error_support.Error{Code: int(rv.FieldByName("Code").Int()), Msg: rv.FieldByName("Msg").String()}
	}
	return nil
}

// SRpcCall rpc请求拓展
func SRpcCall(ctx *routing.Context, serverName, funcName string, args interface{}, reply interface{}) error {
	if rpcClient.clientPool == nil {
		panic("请配置rpcx参数")
	}
	var hc http.HttpController
	if err := hc.ReadJson(ctx, args); err != nil {
		return hc.Fail(ctx, &error_support.Error{Code: 5000, Msg: "数据解析失败"})
	}

	content, _ := jsoniter.MarshalToString(args)
	var traceId string
	if telemetry.GetSpanTraceId() != nil {
		traceId = telemetry.GetSpanTraceId().(string)
	}
	requestHeader := map[string]string{
		"request_id":    stringx.Strval(logger.GetRequestId()),
		"span_trace_id": traceId,
		"client_name":   rpcClient.ClientName,
		"content":       content,
	}
	requestHeader = rpcClient.headerFunc(ctx, requestHeader)
	c := rpcx.Header().Set(context.Background(), requestHeader)
	err := rpcClient.clientPool.Client().Call(c, serverName, funcName, args, reply)
	if err != nil {
		logger.Error("ServerName: ", serverName, ", FuncName: ", funcName, ", Err: ", err)
		return hc.Fail(ctx, &error_support.Error{Code: 5001, Msg: "服务器异常"})
	}

	rv := reflect.ValueOf(reply)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.FieldByName("Code").Int() != int64(0) {
		return hc.Fail(ctx, &error_support.Error{Code: int(rv.FieldByName("Code").Int()), Msg: rv.FieldByName("Msg").String()})
	}

	return hc.Success(ctx, reply, rv.FieldByName("Msg").String())
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
	}, rpcPlugins)
	if err != nil {
		panic(err)
	}
	if len(auth) > 0 {
		rpcServer.RegisterAuthFunc(auth[0])
	}
	rpcServer.Start()
}
