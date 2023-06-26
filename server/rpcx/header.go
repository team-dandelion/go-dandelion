package rpcx

import (
	"context"
	"github.com/smallnest/rpcx/share"
	"strconv"
)

func Header() *RPCHeader {
	return &RPCHeader{}
}

type RPCHeader struct {
}

func (h *RPCHeader) Set(ctx context.Context, header map[string]string) context.Context {
	return context.WithValue(ctx, share.ReqMetaDataKey, header)
}

func (h *RPCHeader) get(ctx context.Context, key string) string {
	data := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	return data[key]
}

func (h *RPCHeader) Value(ctx context.Context, key string) string {
	return h.get(ctx, key)
}

func (h *RPCHeader) Int(ctx context.Context, key string) (int, error) {
	vStr := h.get(ctx, key)
	return strconv.Atoi(vStr)
}

func (h *RPCHeader) IntDefault(ctx context.Context, key string, def int) int {
	vStr := h.get(ctx, key)
	value, err := strconv.Atoi(vStr)
	if err != nil {
		return def
	}
	return value
}

func (h *RPCHeader) Int32(ctx context.Context, key string) (int32, error) {
	vStr := h.get(ctx, key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (h *RPCHeader) Int32Default(ctx context.Context, key string, def int32) int32 {
	vStr := h.get(ctx, key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return def
	}
	return int32(value)
}

func (h *RPCHeader) Int64(ctx context.Context, key string) (int32, error) {
	vStr := h.get(ctx, key)
	value, err := strconv.ParseInt(vStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (h *RPCHeader) Int64Default(ctx context.Context, key string, def int64) int64 {
	vStr := h.get(ctx, key)
	value, err := strconv.ParseInt(vStr, 10, 64)
	if err != nil {
		return def
	}
	return value
}

func (h *RPCHeader) Bool(ctx context.Context, key string) (bool, error) {
	vStr := h.get(ctx, key)
	return strconv.ParseBool(vStr)
}

func (h *RPCHeader) BoolDefault(ctx context.Context, key string, def bool) bool {
	vStr := h.get(ctx, key)
	value, err := strconv.ParseBool(vStr)
	if err != nil {
		return def
	}
	return value
}
