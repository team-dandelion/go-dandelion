package http

import (
	"fmt"
	"github.com/gly-hub/go-dandelion/logger"
	jsoniter "github.com/json-iterator/go"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/rs/xid"
	"net/http"
	"strconv"
	"strings"
)

// middlewareRequestLink 请求链路
func middlewareRequestLink(c *routing.Context) error {
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


func middlewareCustomError(c *routing.Context) error {
	defer func() {
		if err := recover(); err != nil {
			c.Abort()
			c.SetStatusCode(200)
			switch errStr := err.(type) {
			case string:
				p := strings.Split(errStr, "#")
				if len(p) == 3 && p[0] == "CustomError" {
					statusCode, e := strconv.Atoi(p[1])
					if e != nil {
						break
					}
					c.SetStatusCode(statusCode)
					c.Error("服务器异常", http.StatusOK)
					return
				}
			default:
				panic(err)
				fmt.Println(err)
				c.Error("服务器异常", 500)
				return
			}
		}
	}()
	return c.Next()
}
