package http

import (
	"encoding/json"
	error_support "github.com/gly-hub/go-dandelion/error-support"
	"github.com/gly-hub/go-dandelion/logger"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
)

type HttpController struct {
}

type Response struct {
	Code      int32       `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	RequestId interface{} `json:"request_id"`
}

func (hc *HttpController) ReadParams() error {
	return nil
}

func (hc *HttpController) ReadJson(ctx *routing.Context, data interface{}) error {
	err := jsoniter.Unmarshal(ctx.PostBody(), data)
	return err
}

func (hc *HttpController) Success(ctx *routing.Context, data interface{}, msg string) error {
	resp := &Response{
		Code:      2000,
		Msg:       msg,
		Data:      data,
		RequestId: logger.GetRequestId(),
	}

	byteD, _ := json.Marshal(resp)

	ctx.Success("application/json", byteD)
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
	return nil
}
