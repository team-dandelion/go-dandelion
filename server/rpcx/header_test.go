package rpcx

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeader(t *testing.T) {
	tests := []struct {
		name string
		want *RPCHeader
	}{
		{
			want: &RPCHeader{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Header(), "Header()")
		})
	}
}

func TestRPCHeader(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "true"})
	test, err := Header().Bool(ctx, "test")
	assert.Nil(t, err)
	assert.Equal(t, test, true)
}

func TestRPCHeader_BoolDefault(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "true"})
	test := Header().BoolDefault(ctx, "test1", true)
	assert.Equal(t, test, true)
}

func TestRPCHeader_Int(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "1"})
	test, err := Header().Int(ctx, "test")
	assert.Nil(t, err)
	assert.Equal(t, test, 1)
}

func TestRPCHeader_Int32(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "1"})
	test, err := Header().Int32(ctx, "test")
	assert.Nil(t, err)
	assert.Equal(t, test, int32(1))
}

func TestRPCHeader_Int32Default(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "1"})
	test := Header().Int32Default(ctx, "test1", int32(2))
	assert.Equal(t, test, int32(2))
}

func TestRPCHeader_Int64(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "1"})
	test, err := Header().Int64(ctx, "test")
	assert.Nil(t, err)
	assert.Equal(t, test, int64(1))
}

func TestRPCHeader_Int64Default(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "1"})
	test := Header().Int64Default(ctx, "test1", int64(2))
	assert.Equal(t, test, int64(2))
}

func TestRPCHeader_IntDefault(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "1"})
	test := Header().IntDefault(ctx, "test1", 2)
	assert.Equal(t, test, 2)
}

func TestRPCHeader_Set(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "1"})
	assert.NotEmpty(t, ctx)
}

func TestRPCHeader_Value(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "1"})
	test := Header().Value(ctx, "test")
	assert.Equal(t, test, "1")
}

func TestRPCHeader_get(t *testing.T) {
	ctx := Header().Set(context.Background(), map[string]string{"test": "1"})
	test := Header().get(ctx, "test")
	assert.Equal(t, test, "1")
}
