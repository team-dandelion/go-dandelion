package telemetry

import (
	"errors"
	"fmt"
	"github.com/opentracing/opentracing-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
)

var (
	Tracer opentracing.Tracer
	Closer io.Closer
)

func InitTracer(serviceName, agentHostPort string) error {
	if Tracer != nil && Closer != nil {
		return nil
	}

	cfg := jaegerConfig.Configuration{
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: agentHostPort,
		},
		ServiceName: serviceName,
	}

	_tracer, _closer, err := cfg.NewTracer()
	if err != nil {
		fmt.Println("Init GlobalJaegerTracer failed, err : %v", err)
		return err
	}

	Tracer = _tracer
	Closer = _closer
	return nil
}

func getParentSpan(operationName, traceId string, startIfNoParent bool) (span opentracing.Span, err error) {
	if Tracer == nil {
		err = errors.New("jaeger tracing error : Tracer is nil")
		fmt.Println(err)
		return
	}

	parentSpanCtx, err := Tracer.Extract(opentracing.TextMap, opentracing.TextMapCarrier{"UBER-TRACE-ID": traceId})
	if err != nil {
		if startIfNoParent {
			span = Tracer.StartSpan(operationName)
		}
	} else {
		span = Tracer.StartSpan(operationName, opentracing.ChildOf(parentSpanCtx))
	}

	err = nil
	return
}

func StartSpan(operationName, parentSpanTraceId string, startIfNoParent bool) (span opentracing.Span, spanTraceId string, err error) {
	if Tracer == nil || Closer == nil {
		return nil, "", errors.New("jaeger tracing error : Tracer or Closer is nil")
	}
	plainParentSpan, err := getParentSpan(operationName, parentSpanTraceId, startIfNoParent)
	if err != nil || plainParentSpan == nil {
		fmt.Println("No span return")
		return
	}

	m := map[string]string{}
	carrier := opentracing.TextMapCarrier(m)
	err = Tracer.Inject(plainParentSpan.Context(), opentracing.TextMap, carrier)
	if err != nil {
		fmt.Println("jaeger tracing inject error : ", err)
		return
	}

	spanTraceId = carrier["uber-trace-id"]
	span = plainParentSpan
	return
}

func FinishSpan(span opentracing.Span) {
	if span != nil {
		span.Finish()
	}
}

func SpanSetTag(span opentracing.Span, tagname string, tagvalue interface{}) {
	if span != nil {
		span.SetTag(tagname, tagvalue)
	}
}
