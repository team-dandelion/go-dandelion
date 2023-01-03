package http

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/system"
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-dandelion/logger"
	jsoniter "github.com/json-iterator/go"
	"github.com/rs/xid"
	"net/http"
	"regexp"
	"runtime"
)

var ignoreResultPath []string

func LogIgnoreResult(path ...string) {
	ignoreResultPath = append(ignoreResultPath, path...)
}

func checkLogIgnoreResult(path string) bool {
	for _, ignore := range ignoreResultPath {
		reg := regexp.MustCompile(ignore)
		if len(reg.FindAllString(path, -1)) > 0 {
			return true
		}
	}
	return false
}

// middlewareRequestLink 请求链路
func middlewareRequestLink() routing.Handler {
	return func(c *routing.Context) error {
		logger.SetRequestId(xid.New())
		// 打印请求日志
		var data = make(map[string]interface{})
		jsoniter.Unmarshal(c.PostBody(), &data)
		body, _ := jsoniter.MarshalToString(data)

		err := c.Next()

		var code = logger.Blue(fmt.Sprintf("[%v]", c.Response.StatusCode()))
		if err != nil {
			switch err.(type) {
			case routing.HTTPError:
				code = logger.Red(fmt.Sprintf("[%v]", err.(routing.HTTPError).StatusCode()))
			}
		}

		if c.Response.StatusCode() != 200 {
			code = logger.Red(fmt.Sprintf("[%v]", c.Response.StatusCode()))
		}

		var result string
		if !checkLogIgnoreResult(string(c.RequestURI())) {
			result = string(c.Response.Body())
			if len(result) > 500 {
				result = result[:462] + "......" + result[len(result)-38:]
			}
		} else {
			result = "响应结果已忽略"
		}

		logger.Info("ip: %s, method: %s, path: %s, params: %v, result: %v %v", c.RequestCtx.Conn().RemoteAddr().String(), string(c.Method()), c.RequestURI(), body, result, code)
		defer func() {
			logger.DeleteRequestId()
		}()

		return err
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
