package application

import (
	"github.com/team-dandelion/go-dandelion/config"
	"github.com/team-dandelion/go-dandelion/telemetry"
)

func initTracer() {
	if config.Conf.Tracer != nil && config.Conf.Tracer.OpenTrace {
		_ = telemetry.InitTracer(config.Conf.Tracer.TraceName, config.Conf.Tracer.Host)
	}
}
