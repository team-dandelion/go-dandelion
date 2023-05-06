package telemetry

import (
	"github.com/petermattis/goid"
	"sync"
)

var (
	traceIds = sync.Map{}
)

func SetSpanTraceId(ID interface{}) {
	goId := goid.Get()
	traceIds.Store(goId, ID)
}

func GetSpanTraceId() interface{} {
	goId := goid.Get()
	id, _ := traceIds.Load(goId)
	return id
}

func DeleteSpanTraceId() {
	goId := goid.Get()
	traceIds.Delete(goId)
}
