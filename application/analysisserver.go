package application

import (
	"errors"
	routing "github.com/gly-hub/fasthttp-routing"
	"github.com/gly-hub/go-dandelion/config"
	"github.com/gly-hub/go-dandelion/server/analysis"
)

var analysisServer *analysis.Server

func initAnalysisServer() {
	if config.Conf.AnalysisServer == nil || config.Conf.AnalysisServer.Pprof == 0 {
		return
	}

	// 初始化分析器
	analysisServer = analysis.New(analysis.Config{
		Port:        config.Conf.AnalysisServer.Pprof,
		Prometheus:  config.Conf.AnalysisServer.Prometheus,
		ServiceName: config.Conf.AnalysisServer.AnalysisName,
	})

	go func() {
		analysisServer.RunServer()
	}()
}

func HttpPrometheus() (routing.Handler, error) {
	if analysisServer == nil || analysisServer.Prometheus == nil {
		return nil, errors.New("not found prometheus config")
	}
	return analysisServer.Prometheus.HttpMiddleware(), nil
}
