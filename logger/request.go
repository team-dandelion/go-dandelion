package logger

import (
	"github.com/petermattis/goid"
	"sync"
)

var (
	requestIds = sync.Map{}
)

func SetRequestId(ID interface{}) {
	goId := goid.Get()
	requestIds.Store(goId, ID)
}

func GetRequestId() interface{} {
	goId := goid.Get()
	id, _ := requestIds.Load(goId)
	return id
}

func DeleteRequestId() {
	goId := goid.Get()
	requestIds.Delete(goId)
}
