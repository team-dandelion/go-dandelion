package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	loggerx "github.com/team-dandelion/go-dandelion/logger"
	"github.com/team-dandelion/go-dandelion/telemetry"
	"go.mongodb.org/mongo-driver/event"
)

type LogLevel int

const (
	Silent LogLevel = iota + 1
	Error
	Warn
	Info
)

type Logger struct {
	Level         LogLevel
	SlowThreshold time.Duration
}

func NewLogger(level LogLevel, slowThreshold time.Duration) *Logger {
	return &Logger{
		Level:         level,
		SlowThreshold: slowThreshold,
	}
}

// CreateMonitor 创建MongoDB监控器，用于记录操作日志
func (logger *Logger) CreateMonitor() (*event.CommandMonitor, *event.PoolMonitor) {
	commandMonitor := &event.CommandMonitor{
		Started: func(ctx context.Context, evt *event.CommandStartedEvent) {
			if logger.Level <= Silent {
				return
			}

			traceId := telemetry.GetSpanTraceId()
			if traceId != nil {
				span, _, _ := telemetry.StartSpan("MongoDB", traceId.(string), false, opentracing.StartTime(time.Now()))
				telemetry.SpanSetTag(span, "operation", evt.CommandName)
				telemetry.SpanSetTag(span, "database", evt.DatabaseName)
				telemetry.SpanSetTag(span, "request_id", loggerx.GetRequestId())
				ctx = context.WithValue(ctx, "mongo_span", span)
			}
		},
		Succeeded: func(ctx context.Context, evt *event.CommandSucceededEvent) {
			if logger.Level <= Silent {
				return
			}

			duration := time.Duration(evt.DurationNanos)

			if span, ok := ctx.Value("mongo_span").(opentracing.Span); ok {
				span.LogFields(log.String("status", "success"))
				span.LogFields(log.String("duration", duration.String()))
				telemetry.FinishSpan(span)
			}

			logger.logOperation(evt.CommandName, evt.DatabaseName, "", duration, nil)
		},
		Failed: func(ctx context.Context, evt *event.CommandFailedEvent) {
			if logger.Level < Error {
				return
			}

			duration := time.Duration(evt.DurationNanos)

			if span, ok := ctx.Value("mongo_span").(opentracing.Span); ok {
				span.LogFields(log.String("status", "failed"))
				span.LogFields(log.String("error", evt.Failure))
				span.LogFields(log.String("duration", duration.String()))
				telemetry.FinishSpan(span)
			}

			logger.logOperation(evt.CommandName, evt.DatabaseName, evt.Failure, duration, fmt.Errorf(evt.Failure))
		},
	}

	poolMonitor := &event.PoolMonitor{
		Event: func(evt *event.PoolEvent) {
			if logger.Level >= Info {
				switch evt.Type {
				case "ConnectionPoolCreated":
					loggerx.Info(fmt.Sprintf("%v MongoDB连接池已创建 - 地址: %v",
						loggerx.Blue("[mongodb]"), evt.Address))
				case "ConnectionPoolClosed":
					loggerx.Info(fmt.Sprintf("%v MongoDB连接池已关闭 - 地址: %v",
						loggerx.Blue("[mongodb]"), evt.Address))
				case "ConnectionCreated":
					loggerx.Info(fmt.Sprintf("%v 新建MongoDB连接 - 地址: %v",
						loggerx.Blue("[mongodb]"), evt.Address))
				case "ConnectionClosed":
					loggerx.Info(fmt.Sprintf("%v MongoDB连接已关闭 - 地址: %v",
						loggerx.Blue("[mongodb]"), evt.Address))
				}
			}
		},
	}

	return commandMonitor, poolMonitor
}

func (logger *Logger) logOperation(operation, database, failure string, duration time.Duration, err error) {
	durationMs := float64(duration.Nanoseconds()) / 1e6

	switch {
	case err != nil && logger.Level >= Error:
		msg := fmt.Sprintf("%v %v %v %v",
			loggerx.Blue(fmt.Sprintf("[mongodb] [%.3fms]", durationMs)),
			fmt.Sprintf("操作: %s, 数据库: %s", operation, database),
			loggerx.Red(failure),
		)
		loggerx.Error(msg)

	case duration > logger.SlowThreshold && logger.SlowThreshold > 0 && logger.Level >= Warn:
		msg := fmt.Sprintf("%v %v %v",
			loggerx.Blue("[mongodb]"),
			loggerx.Red(fmt.Sprintf("[%.3fms]>= %v", durationMs, logger.SlowThreshold)),
			fmt.Sprintf("慢操作: %s, 数据库: %s", operation, database))
		loggerx.Warn(msg)

	case logger.Level >= Info:
		msg := fmt.Sprintf("%v %v",
			loggerx.Blue(fmt.Sprintf("[mongodb] [%.3fms]", durationMs)),
			fmt.Sprintf("操作: %s, 数据库: %s", operation, database))
		loggerx.Info(msg)
	}
}

func (logger *Logger) Info(message string, args ...interface{}) {
	if logger.Level >= Info {
		loggerx.Info(message, args...)
	}
}

func (logger *Logger) Warn(message string, args ...interface{}) {
	if logger.Level >= Warn {
		loggerx.Warn(message, args...)
	}
}

func (logger *Logger) Error(message string, args ...interface{}) {
	if logger.Level >= Error {
		loggerx.Error(message, args...)
	}
}
