package http

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/system"
	"github.com/gly-hub/go-dandelion/logger"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/rs/xid"
	"net/http"
	"runtime"
)

// middlewareRequestLink 请求链路
func middlewareRequestLink() routing.Handler {
	return func(c *routing.Context) error {
		logger.SetRequestId(xid.New())
		// 打印请求日志
		//body := strings.ReplaceAll(string(), "\n", "")
		//body = strings.ReplaceAll(body, "\t", "")
		var data = make(map[string]interface{})
		jsoniter.Unmarshal(c.PostBody(), &data)
		body, _ := jsoniter.MarshalToString(data)

		logger.Info("ip: %s, path: %s, params: %v", c.RequestCtx.Conn().RemoteAddr().String(), c.RequestURI(), body)
		defer func() {
			logger.DeleteRequestId()
		}()

		return c.Next()
	}
}

func middlewareCustomError() routing.Handler {
	return func(c *routing.Context) error {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 4096)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				err = fmt.Errorf("[service internal error]: %v, stack: %s",
					err, buf)
				logger.Error(err)
			}
		}()
		return c.Next()
	}
}

func middlewareSentinel() routing.Handler {
	if _, err := system.LoadRules([]*system.Rule{
		{
			MetricType:   system.InboundQPS,
			TriggerCount: 200,
			Strategy:     system.BBR,
		},
	}); err != nil {
		logger.Error("Unexpected error: %+v", err)
	}

	return sentinelFunc(WithBlockFallback(func(ctx *routing.Context) error {
		ctx.SetStatusCode(500)
		ctx.Error("too many request; the quota used up", 500)
		return nil
	}))
}

func sentinelFunc(opts ...Option) routing.Handler {
	options := evaluateOptions(opts)
	return func(ctx *routing.Context) error {
		resourceName := string(ctx.Method()) + ":" + string(ctx.Path())

		if options.resourceExtract != nil {
			resourceName = options.resourceExtract(ctx)
		}

		entry, entryErr := sentinel.Entry(
			resourceName,
			sentinel.WithResourceType(base.ResTypeWeb),
			sentinel.WithTrafficType(base.Inbound),
		)

		if entryErr != nil {
			if options.blockFallback != nil {
				return options.blockFallback(ctx)
			} else {
				ctx.SetStatusCode(http.StatusTooManyRequests)
				return nil
			}
		}

		defer entry.Exit()
		return ctx.Next()
	}
}
