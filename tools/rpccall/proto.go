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

func SProtoCall(ctx *routing.Context, param interface{}, handler interface{}) error {
	rpcClient := application.GetRpcClient()
	if rpcClient.ClientPool == nil {
		panic("请配置rpcx参数")
	}

	var hc http.HttpController
	refParam := reflect.ValueOf(param)
	if refParam.Kind() != reflect.Ptr {
		return hc.Fail(ctx, &error_support.Error{Code: 5000, Msg: "数据解析失败"})
	}

	if len(ctx.PostBody()) > 0 {
		if err := hc.ReadJson(ctx, param); err != nil {
			return hc.Fail(ctx, &error_support.Error{Code: 5000, Msg: "数据解析失败"})
		}

	}

	content, _ := jsoniter.MarshalToString(param)
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
	newCtx := rpcx.Header().Set(context.Background(), requestHeader)
	refHandler := reflect.ValueOf(handler)
	if refHandler.Kind() != reflect.Func {
		return hc.Fail(ctx, &error_support.Error{Code: 5000, Msg: "WRONG handler"})
	}

	inParam := []reflect.Value{reflect.ValueOf(newCtx), refParam}
	rets := refHandler.Call(inParam)
	if !rets[1].IsNil() {
		err := rets[1].Interface().(error)
		return hc.Fail(ctx, &error_support.Error{Code: 5000, Msg: err.Error()})
	}
	// 对rpc响应内容进行处理
	rt := reflect.TypeOf(rets[0].Interface())
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	_, cOk := rt.FieldByName("CommonResp")
	if cOk {
		rt2 := reflect.ValueOf(rets[0].Interface())
		if rt2.Kind() == reflect.Ptr {
			rt2 = rt2.Elem()
		}

		rv := rt2.FieldByName("CommonResp").Elem()
		if !rv.IsValid() {
			return hc.Success(ctx, rets[0].Interface(), "")
		}
		if rv.FieldByName("Code").IsValid() && rv.FieldByName("Code").Int() != int64(0) {
			return hc.Fail(ctx, &error_support.Error{Code: int(rv.FieldByName("Code").Int()), Msg: rv.FieldByName("Msg").String()})
		}
		return hc.Success(ctx, rets[0].Interface(), rv.FieldByName("Msg").String())
	}

	return hc.Success(ctx, rets[0].Interface(), "")
}
