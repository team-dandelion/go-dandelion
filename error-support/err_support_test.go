package error_support

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	out := &struct {
		Code int32  `json:"code"`
		Msg  string `json:"msg"`
	}{}
	Format(&Error{
		Code: 2000,
		Msg:  "test",
	}, out)
	assert.Equal(t, int32(2000), out.Code)
	assert.Equal(t, "test", out.Msg)
}
