package telemetry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetSpanTraceId(t *testing.T) {
	SetSpanTraceId("test")
	id := GetSpanTraceId()
	assert.Equal(t, id, "test")
	DeleteSpanTraceId()
}
