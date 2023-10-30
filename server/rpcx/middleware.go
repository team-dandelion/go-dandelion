package rpcx

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/petermattis/goid"
	"github.com/smallnest/rpcx/protocol"
	"github.com/team-dandelion/go-dandelion/logger"
	"github.com/team-dandelion/go-dandelion/telemetry"
	"sync"
	"time"
)

var (
	spanMap sync.Map
)

type ClientLoggerPlugin struct {
}

func (p *ClientLoggerPlugin) PostCall(ctx context.Context, servicePath, serviceMethod string, args interface{}) error {
	requestId := logger.GetRequestId()
	ctx = context.WithValue(ctx, "request_id", requestId)
	return nil
}

type ServerLoggerPlugin struct {
}

func (p *ServerLoggerPlugin) PreHandleRequest(ctx context.Context, r *protocol.Message) error {
	logger.SetRequestId(r.Metadata["request_id"])
	traceId := r.Metadata["span_trace_id"]
	if traceId != "" {
		span, spanTraceId, err := telemetry.StartSpan("RpcCall", traceId, true, opentracing.StartTime(time.Now()))
		if err == nil {
			telemetry.SpanSetTag(span, "request_id", r.Metadata["request_id"])
			telemetry.SpanSetTag(span, "call_method", r.ServiceMethod)
			telemetry.SetSpanTraceId(spanTraceId)
			spanMap.Store(goid.Get(), span)
		}
	}

	logger.Info("client: %s, server: %v, func: %s, params: %s", r.Metadata["client_name"], r.ServicePath, r.ServiceMethod, r.Metadata["content"])
	return nil
}

func (p *ServerLoggerPlugin) PostWriteResponse(ctx context.Context, req *protocol.Message, res *protocol.Message, err error) error {
	logger.DeleteRequestId()
	if span, ok := spanMap.Load(goid.Get()); ok {
		telemetry.FinishSpan(span.(opentracing.Span))
		spanMap.Delete(goid.Get())
	}
	return nil
}
