package logger

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type fileWrite struct {
	//sync.Mutex
}

func newFile() logger {
	fw := &fileWrite{}
	return fw
}

func (f *fileWrite) init() error {
	return nil
}

func (f *fileWrite) writeMsg(when time.Time, msg string, level int) error {
	//	f.Lock()
	//	defer f.Unlock()
	b, _ := formatTimeHeader(when)
	f.writeFile(string(b) + " " + msg)
	return nil
}

func (f *fileWrite) destroy() {

}
func (f *fileWrite) flush() {

}

func (f *fileWrite) writeFile(msg string) {
	var fi *os.File
	var err error
	createDir("logs")
	var _fileName string = "logs/" + time.Now().Format("20060102") + ".log"
	_exist, _size := checkFile(_fileName)
	if _exist && _size > 2048*1024 {
		errr := os.Rename(_fileName, "logs/"+time.Now().Format("20060102150405")+".log")
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
	//	if err == nil {
	//		_, err = io.WriteString(fi, msg+"\r")
	//		fi.Close()
	//	}

	fi, err = os.OpenFile(_fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer fi.Close()
	if err == nil {
		w := bufio.NewWriter(fi)
		w.WriteString(msg + "\r")
		w.Flush()
	}
}
