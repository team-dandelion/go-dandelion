package rpccall

import (
	"context"
	"reflect"

	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/toolbox/stringx"
	jsoniter "github.com/json-iterator/go"
	"github.com/team-dandelion/go-dandelion/application"
	error_support "github.com/team-dandelion/go-dandelion/error-support"
	"github.com/team-dandelion/go-dandelion/logger"
	"github.com/team-dandelion/go-dandelion/server/rpcx"
	"github.com/team-dandelion/go-dandelion/telemetry"
)

func SProtoCall(ctx *routing.Context, param interface{}, handler interface{}) error {
	rpcClient := application.GetRpcClient()
	if rpcClient.ClientPool == nil {
		panic("请配置rpcx参数")
	}

	var hc HttpController
	refParam := reflect.ValueOf(param)
	if refParam.Kind() != reflect.Ptr {
		return hc.Fail(ctx, &error_support.Error{Code: RPCServerFailCode, Msg: "数据解析失败"})
	}

	if string(ctx.Request.Header.ContentType()) == "application/json" && len(ctx.PostBody()) > 0 {
		logger.Info(string(ctx.PostBody()))
		if err := hc.ReadJson(ctx, param); err != nil {
			return hc.Fail(ctx, &error_support.Error{Code: RPCServerFailCode, Msg: "数据解析失败"})
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
		return hc.Fail(ctx, &error_support.Error{Code: SystemFailCode, Msg: "WRONG handler"})
	}

	inParam := []reflect.Value{reflect.ValueOf(newCtx), refParam}
	rets := refHandler.Call(inParam)
	if !rets[1].IsNil() {
		errValue := rets[1].Interface().(error)
		return hc.Fail(ctx, &error_support.Error{Code: RPCServerFailCode, Msg: errValue.Error()})
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

func ProtoCall(ctx *routing.Context, handler interface{}, param interface{},
	resp interface{}) error {
	rpcClient := application.GetRpcClient()
	if rpcClient.ClientPool == nil {
		panic("请配置rpcx参数")
	}

	refParam := reflect.ValueOf(param)
	if refParam.Kind() != reflect.Ptr {
		return &error_support.Error{Code: 5000, Msg: "数据解析失败"}
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
		return &error_support.Error{Code: 5000, Msg: "WRONG handler"}
	}

	inParam := []reflect.Value{reflect.ValueOf(newCtx), refParam}
	rets := refHandler.Call(inParam)
	if !rets[1].IsNil() {
		err := rets[1].Interface().(error)
		logger.Error(err)
		return err
	}

	refResp := reflect.ValueOf(resp)
	if refResp.Kind() != reflect.Ptr {
		return &error_support.Error{Code: 5000, Msg: "数据解析失败"}
	}

	refResp.Elem().Set(rets[0].Elem())
	return nil
}

func CommonRespFormat(ret interface{}, bErr error) error {
	var err error
	defer func() {
		if err1 := recover(); err1 != nil {
			err = err1.(error)
		}
	}()

	rt := reflect.TypeOf(ret)
	var kind reflect.Kind
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
		kind = rt.Kind()
	}
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}
	inst := reflect.New(rt)
	var commonResp = reflect.ValueOf(&CommonResp{
		Code: SuccessCode,
		Msg:  "",
	})

	if bErr != nil {
		var (
			code int
			msg  string
		)
		switch bErr.(type) {
		case *error_support.Error:
			code = bErr.(*error_support.Error).Code
			msg = bErr.(*error_support.Error).Error()
		default:
			code = SystemFailCode
			msg = bErr.Error()
		}
		commonResp = reflect.ValueOf(&CommonResp{
			Code: int32(code),
			Msg:  msg,
		})
	}

	inst.Elem().FieldByName("CommonResp").Set(commonResp)
	if kind == reflect.Ptr {
		reflect.ValueOf(ret).Elem().Set(inst)
	} else {
		reflect.ValueOf(ret).Elem().Set(inst.Elem())
	}

	return err
}
