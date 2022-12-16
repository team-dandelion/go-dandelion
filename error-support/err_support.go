package error_support

import (
	"io/ioutil"
	"reflect"
)

var errCodeMap map[int]*Error

type Error struct {
	Name string
	Code int
	Msg  string
}

func (e *Error) Error() string {
	return e.Msg
}

func Init(path string) {
	errCodeMap = make(map[int]*Error)
	showFileList(path)
	for _, route := range routes {
		bt, _ := ioutil.ReadFile(route)
		src := string(bt)
		result := scanFuncDeclByComment(``, src, "@ErrMsg[(].*?[)]")
		for k, v := range result.Errors() {
			errCodeMap[k] = v
		}
	}
}

func Format(err error, out interface{}) {
	var code int32
	var msg string
	switch err.(type) {
	case *Error:
		code = int32(err.(*Error).Code)
		msg = err.(*Error).Msg
		if _, ok := errCodeMap[err.(*Error).Code]; ok {
			msg = errCodeMap[err.(*Error).Code].Msg
		}
	default:
		code = 40000
		msg = err.Error()
	}

	// 尝试获取common_resp
	rv := reflect.ValueOf(out)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.FieldByName("Code").CanAddr() {
		rv.FieldByName("Code").SetInt(int64(code))
	}

	if rv.FieldByName("Msg").CanAddr() {
		rv.FieldByName("Msg").SetString(msg)
	}
}
