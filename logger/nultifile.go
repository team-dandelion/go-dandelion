package logger

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

type multiFileWrite struct {
	sync.Mutex
}

func newMultiFile() logger {
	fw := &multiFileWrite{}
	return fw
}

func (f *multiFileWrite) init() error {
	return nil
}

func (f *multiFileWrite) writeMsg(when time.Time, msg string, level int) error {
	//	f.Lock()
	//	defer f.Unlock()
	b, _ := formatTimeHeader(when)
	f.writeFile(string(b)+" "+msg, level)
	return nil
}

func (f *multiFileWrite) destroy() {

}
func (f *multiFileWrite) flush() {

}

func (f *multiFileWrite) writeFile(msg string, level int) {
	var fi *os.File
	var err error
	createDir("logs")
	var _fileName string = "logs/" + levelName[level] + time.Now().Format("20060102") + ".log"
	_exist, _size := checkFile(_fileName)
	if _exist && _size > 2048*1024 {
		errr := os.Rename(_fileName, "logs/"+levelName[level]+time.Now().Format("20060102150405")+".log")
		if errr != nil {
			fmt.Println(errr.Error())
		}
		_exist = false
	}
	//	if _exist {
	//		fi, err = os.OpenFile(_fileName, os.O_APPEND, 0666)
	//	} else {
	//		fi, err = os.Create(_fileName)
	//	}
	//if err == nil {
	//_, err = io.WriteString(fi, msg+"\r")
	//fi.Close()
	//}
	//以上这种方式在MAC系统下不能正确追加

	fi, err = os.OpenFile(_fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer fi.Close()
	if err == nil {
		w := bufio.NewWriter(fi)
		w.WriteString(msg + "\r")
		w.Flush()
	}
}
