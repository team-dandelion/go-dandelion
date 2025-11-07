package rpccall

import (
	"encoding/json"

	routing "github.com/gly-hub/fasthttp-routing"
	jsoniter "github.com/json-iterator/go"
	error_support "github.com/team-dandelion/go-dandelion/error-support"
	"github.com/team-dandelion/go-dandelion/logger"
)

const (
	SuccessCode          = 20000
	SystemFailCode       = 50000
	ParamsFormatFailCode = 50001 // 参数解析错误
	RPCServerFailCode    = 50002 // RPC服务异常
)

type HttpController struct {
}

type Response struct {
	Code      int32       `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	RequestId interface{} `json:"requestId"`
}

func (hc *HttpController) ReadParams() error {
	return nil
}

func (hc *HttpController) ReadJson(ctx *routing.Context, data interface{}) error {
	if len(ctx.PostBody()) == 0 {
		return nil
	}
	err := jsoniter.Unmarshal(ctx.PostBody(), data)
	return err
}

func (hc *HttpController) Success(ctx *routing.Context, data interface{}, msg string) error {
	resp := &Response{
		Code:      SuccessCode,
		Msg:       msg,
		Data:      data,
		RequestId: logger.GetRequestId(),
	}

	byteD, _ := json.Marshal(resp)

	ctx.Success("application/json", byteD)
	ctx.Abort()
	return nil
}

func (hc *HttpController) Fail(ctx *routing.Context, err error, msg ...string) error {
	resp := &Response{
		RequestId: logger.GetRequestId(),
	}

	error_support.Format(err, resp)

	if len(msg) > 0 {
		resp.Msg = msg[0]
	}

	byteD, _ := json.Marshal(resp)

	ctx.Success("application/json", byteD)
	ctx.Abort()
	return nil
}
