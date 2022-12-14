package logger

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type consoleWrite struct {
	sync.Mutex
	colorful bool
}

func newConsole() logger {
	cw := &consoleWrite{
		colorful: runtime.GOOS != "windows",
	}
	return cw
}

func (c *consoleWrite) init() error {
	if runtime.GOOS == "windows" {
		c.colorful = false
	} else {
		c.colorful = true
	}
	return nil
}

func (c *consoleWrite) writeMsg(when time.Time, msg string, level int) error {
	c.Lock()
	defer c.Unlock()
	b, _ := formatTimeHeader(when)

	//Level 0 紧急的 1警报 2重要的 3错误 4警告 5提示 6信息 7调试
	//fmt.Println(c.colorful)
	//windows字体有问题 采用git bash 启动可实现
	if !c.colorful {
		fmt.Println(string(b), msg)
	} else {
		switch level {
		case 6:
			fmt.Printf("%v %v \n", string(b), Green(msg))
		case 3:
			fmt.Printf("%v %v \n", string(b), Red(msg))
		case 4:
			fmt.Printf("%v %v \n", string(b), Yellow(msg))
		case 7:
			fmt.Printf("%v %v \n", string(b), Blue(msg))
		}
	}
	return nil
}
func (c *consoleWrite) destroy() {

}
func (c *consoleWrite) flush() {

}
