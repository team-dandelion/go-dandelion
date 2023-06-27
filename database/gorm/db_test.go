package gorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConnection(t *testing.T) {
	db := NewConnection(&Config{
		DBType:        "mysql",
		MaxOpenConn:   10,
		MaxIdleConn:   10,
		MaxLifeTime:   10,
		MaxIdleTime:   10,
		Level:         4,
		SlowThreshold: 10,
		Master: &Master{
			User:     "root",
			Password: "starunion",
			Host:     "127.0.0.1",
			Port:     "6379",
			Database: "go-admin",
		},
	})
	assert.NotEmpty(t, db)
}
