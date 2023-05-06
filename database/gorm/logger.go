package gorm

import (
	"context"
	"errors"
	"fmt"
	loggerx "github.com/gly-hub/go-dandelion/logger"
	"github.com/gly-hub/go-dandelion/telemetry"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"strconv"
	"time"
)

type Logger struct {
	Level                     glogger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func (logger *Logger) LogMode(level glogger.LogLevel) glogger.Interface {
	logger.Level = level
	return logger
}

func (logger *Logger) Info(ctx context.Context, s string, i ...interface{}) {
	if logger.Level >= glogger.Info {
		loggerx.Info(s, i)
	}
}

func (logger *Logger) Warn(ctx context.Context, s string, i ...interface{}) {
	if logger.Level >= glogger.Warn {
		loggerx.Warn(s, i)
	}
}

func (logger *Logger) Error(ctx context.Context, s string, i ...interface{}) {
	if logger.Level >= glogger.Error {
		loggerx.Error(s, i)
	}
}

func (logger *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if logger.Level <= glogger.Silent {
		return
	}

	tranceId := telemetry.GetSpanTraceId()
	var strSql string
	elapsed := time.Since(begin)
	if tranceId != nil {
		span, _, _ := telemetry.StartSpan("GORM", tranceId.(string), false, opentracing.StartTime(begin))
		telemetry.SpanSetTag(span, "request_id", loggerx.GetRequestId())
		defer func() {
			span.LogFields(log.String("sql", strSql))
			span.LogFields(log.String("elapsed", elapsed.String()))
			telemetry.FinishSpan(span)
		}()
	}
	switch {
	case err != nil && logger.Level >= glogger.Error && (!errors.Is(err, glogger.ErrRecordNotFound) || !logger.IgnoreRecordNotFoundError):
		sql, rows := fc()
		rowStr := ""
		if rows == -1 {
			rowStr = "-"
		} else {
			rowStr = strconv.FormatInt(rows, 10)
		}
		msg := fmt.Sprintf("%v %v %v\t%v",
			loggerx.Blue(fmt.Sprintf("[gorm] [%.3fms] [rows:%v]", float64(elapsed.Nanoseconds())/1e6, rowStr)),
			sql,
			loggerx.Red(err.Error()),
			utils.FileWithLineNum())
		strSql = sql
		loggerx.Info(msg)
	case elapsed > logger.SlowThreshold && logger.SlowThreshold != 0 && logger.Level >= glogger.Warn:
		sql, rows := fc()
		rowStr := ""
		if rows == -1 {
			rowStr = "-"
		} else {
			rowStr = strconv.FormatInt(rows, 10)
		}
		msg := fmt.Sprintf("%v %v %v %v\t%v",
			loggerx.Blue("[gorm]"),
			loggerx.Red(fmt.Sprintf("[%.3fms]>= %v", float64(elapsed.Nanoseconds())/1e6, logger.SlowThreshold)),
			loggerx.Blue(fmt.Sprintf("[rows:%v]", rowStr)),
			sql,
			utils.FileWithLineNum())
		strSql = sql
		loggerx.Info(msg)
	case logger.Level == glogger.Info:
		sql, rows := fc()
		rowStr := ""
		if rows == -1 {
			rowStr = "-"
		} else {
			rowStr = strconv.FormatInt(rows, 10)
		}
		msg := fmt.Sprintf("%v %v",
			loggerx.Blue(fmt.Sprintf("[gorm] [%.3fms] [rows:%v]", float64(elapsed.Nanoseconds())/1e6, rowStr)),
			sql)
		strSql = sql
		loggerx.Info(msg)
	}
}
