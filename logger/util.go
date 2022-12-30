package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const rfc3339Milli = "2006-01-02 15:04:05.000" //999后面是0则省略
const rfc3339Micro = "2006-01-02 15:04:05.000000"
const rfc3339Nano = "2006-01-02 15:04:05.999999999"

func FormatTimeHeader(when time.Time) ([]byte, int) {
	return formatTimeHeader(when)
}

func formatTimeHeader(when time.Time) ([]byte, int) {
	//var str string = when.Format("2006-01-02 15:04:05")
	var str string = when.Format(rfc3339Milli)
	str += " "
	b := []byte(str)
	return b, len(b)
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}

func checkFile(fileName string) (exist bool, size int64) {
	exist = true
	size = 0
	var fileInfo os.FileInfo
	var err error
	if fileInfo, err = os.Stat(fileName); os.IsNotExist(err) {
		exist = false
	}
	if exist {
		size = fileInfo.Size()
	}
	return
}

func createDir(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.MkdirAll(dirName, 0777)
	}
}
