package rpcx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_makeLocalName(t *testing.T) {
	assert.Equal(t, makeLocalName("test", 1), "test#1")
}

func Test_nodeId(t *testing.T) {
	assert.NotEqual(t, nodeId("127.0.0.1"), 0)
}

func Test_parsePort(t *testing.T) {
	assert.Equal(t, parsePort("127.0.0.1:8080"), ":8080")
}

func Test_serverAddress(t *testing.T) {
	assert.Equal(t, serverAddress("127.0.0.1"), "tcp@127.0.0.1")
}
