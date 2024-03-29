package rpccall

import (
	"context"
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/toolbox/stringx"
	jsoniter "github.com/json-iterator/go"
	"github.com/team-dandelion/go-dandelion/application"
	error_support "github.com/team-dandelion/go-dandelion/error-support"
	"github.com/team-dandelion/go-dandelion/logger"
	"github.com/team-dandelion/go-dandelion/server/http"
	"github.com/team-dandelion/go-dandelion/server/rpcx"
	"github.com/team-dandelion/go-dandelion/telemetry"
	"reflect"
)

// RpcCall rpc请求
func RpcCall(ctx *routing.Context, serverName, funcName string, args interface{}, reply interface{}) error {
	rpcClient := application.GetRpcClient()

	if rpcClient.ClientPool == nil {
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

	requestHeader = rpcClient.HeaderFunc(ctx, requestHeader)
	c := rpcx.Header().Set(context.Background(), requestHeader)
	err := rpcClient.ClientPool.Client().Call(c, serverName, funcName, args, reply)
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
	rpcClient := application.GetRpcClient()
	if rpcClient.ClientPool == nil {
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
	requestHeader = rpcClient.HeaderFunc(ctx, requestHeader)
	c := rpcx.Header().Set(context.Background(), requestHeader)
	err := rpcClient.ClientPool.Client().Call(c, serverName, funcName, args, reply)
	if err != nil {
		logger.Error("ServerName: ", serverName, ", FuncName: ", funcName, ", Err: ", err)
		return hc.Fail(ctx, &error_support.Error{Code: 5001, Msg: "服务器异常"})
	}

	rt := reflect.TypeOf(reply)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	_, cOk := rt.FieldByName("Code")
	_, mOk := rt.FieldByName("Msg")
	if mOk && cOk {
		rv := reflect.ValueOf(reply)
		if rv.Kind() == reflect.Ptr {
			rv = rv.Elem()
		}

		if rv.FieldByName("Code").Int() != int64(0) {
			return hc.Fail(ctx, &error_support.Error{Code: int(rv.FieldByName("Code").Int()), Msg: rv.FieldByName("Msg").String()})
		}
		return hc.Success(ctx, reply, rv.FieldByName("Msg").String())
	}

	return hc.Success(ctx, reply, "")
}
