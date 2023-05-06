package logger

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/telemetry"
	"github.com/gomodule/redigo/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"time"
)

type LoggingConn struct {
	redis.Conn
}

func (c *LoggingConn) Close() error {
	err := c.Conn.Close()
	return err
}

func (c *LoggingConn) printValue(buf *bytes.Buffer, v interface{}) {
	const chop = 32
	switch v := v.(type) {
	case []byte:
		if len(v) > chop {
			fmt.Fprintf(buf, "%q...", v[:chop])
		} else {
			fmt.Fprintf(buf, "%q", v)
		}
	case string:
		if len(v) > chop {
			fmt.Fprintf(buf, "%q...", v[:chop])
		} else {
			fmt.Fprintf(buf, "%q", v)
		}
	case []interface{}:
		if len(v) == 0 {
			buf.WriteString("[]")
		} else {
			sep := "["
			fin := "]"
			if len(v) > chop {
				v = v[:chop]
				fin = "...]"
			}
			for _, vv := range v {
				buf.WriteString(sep)
				c.printValue(buf, vv)
				sep = ", "
			}
			buf.WriteString(fin)
		}
	default:
		fmt.Fprint(buf, v)
	}
}

func (c *LoggingConn) print(method, commandName string, args []interface{}, reply interface{}, err error, startTime time.Time) {
	var buf bytes.Buffer
	tranceId := telemetry.GetSpanTraceId()
	if tranceId != nil {
		span, _, _ := telemetry.StartSpan("Redis", tranceId.(string), false, opentracing.StartTime(startTime))
		telemetry.SpanSetTag(span, "method", method)
		telemetry.SpanSetTag(span, "request_id", logger.GetRequestId())
		defer func() {
			span.LogFields(log.String(method, buf.String()))
			telemetry.FinishSpan(span)
		}()
	}
	fmt.Fprintf(&buf, "%s%s(", logger.Blue("[redigo] "), method)
	if method != "Receive" {
		if commandName == "" {
			return
		}
		buf.WriteString(commandName)
		for _, arg := range args {
			buf.WriteString(", ")
			c.printValue(&buf, arg)
		}
	}
	buf.WriteString(") -> (")
	if method != "Send" {
		c.printValue(&buf, reply)
		buf.WriteString(", ")
	}
	fmt.Fprintf(&buf, "%v)", err)
	logger.Info(buf.String()) // nolint: errcheck
}

func (c *LoggingConn) Do(commandName string, args ...interface{}) (interface{}, error) {
	startTime := time.Now()
	reply, err := c.Conn.Do(commandName, args...)
	c.print("Do", commandName, args, reply, err, startTime)
	return reply, err
}

func (c *LoggingConn) DoContext(ctx context.Context, commandName string, args ...interface{}) (interface{}, error) {
	startTime := time.Now()
	reply, err := redis.DoContext(c.Conn, ctx, commandName, args...)
	c.print("DoContext", commandName, args, reply, err, startTime)
	return reply, err
}

func (c *LoggingConn) DoWithTimeout(timeout time.Duration, commandName string, args ...interface{}) (interface{}, error) {
	startTime := time.Now()
	reply, err := redis.DoWithTimeout(c.Conn, timeout, commandName, args...)
	c.print("DoWithTimeout", commandName, args, reply, err, startTime)
	return reply, err
}

func (c *LoggingConn) Send(commandName string, args ...interface{}) error {
	startTime := time.Now()
	err := c.Conn.Send(commandName, args...)
	c.print("Send", commandName, args, nil, err, startTime)
	return err
}

func (c *LoggingConn) Receive() (interface{}, error) {
	startTime := time.Now()
	reply, err := c.Conn.Receive()
	c.print("Receive", "", nil, reply, err, startTime)
	return reply, err
}

func (c *LoggingConn) ReceiveContext(ctx context.Context) (interface{}, error) {
	startTime := time.Now()
	reply, err := redis.ReceiveContext(c.Conn, ctx)
	c.print("ReceiveContext", "", nil, reply, err, startTime)
	return reply, err
}

func (c *LoggingConn) ReceiveWithTimeout(timeout time.Duration) (interface{}, error) {
	startTime := time.Now()
	reply, err := redis.ReceiveWithTimeout(c.Conn, timeout)
	c.print("ReceiveWithTimeout", "", nil, reply, err, startTime)
	return reply, err
}
