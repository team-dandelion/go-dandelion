package logger

import (
	"sync"
	"github.com/petermattis/goid"
)

var (
	requestIds = sync.Map{}
)

func SetRequestId(ID interface{}){
	goId := goid.Get()
	requestIds.Store(goId, ID)
}

func GetRequestId() interface{}{
	goId := goid.Get()
	id, _ := requestIds.Load(goId)
	return id
}

func DeleteRequestId(){
	goId := goid.Get()
	requestIds.Delete(goId)
}
