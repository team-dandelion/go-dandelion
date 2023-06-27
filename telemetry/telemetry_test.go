package telemetry

import (
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInitTracer(t *testing.T) {
	assert.Emptyf(t, InitTracer("test", ""), "InitTracer should be empty")
}

func TestStartSpan(t *testing.T) {
	InitTracer("test", "")
	gotSpan, gotSpanTraceId, _ := StartSpan("TestServer", "test1", false, opentracing.StartTime(time.Now()))
	assert.Equal(t, gotSpan, opentracing.Span(nil))
	assert.Equal(t, gotSpanTraceId, "")
}

func TestFinishSpan(t *testing.T) {
	FinishSpan(opentracing.Span(nil))
}

func TestSpanSetTag(t *testing.T) {
	SpanSetTag(opentracing.Span(nil), "test", "test")
}
