package rpcx

import (
	"context"
	"github.com/smallnest/rpcx/share"
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
