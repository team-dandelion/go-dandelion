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

	var infoColor, ErrorColor, WarnColor, DebugColor string
	//Level 0 紧急的 1警报 2重要的 3错误 4警告 5提示 6信息 7调试
	//fmt.Println(c.colorful)
	//windows字体有问题 采用git bash 启动可实现
	if !c.colorful {
		fmt.Println(string(b), msg)
	} else {
		switch level {
		case 6:
			infoColor = fmt.Sprintf("%c[%d;%d;%dm%s", 0x1B, 1, 40, 37, "")
			fmt.Printf("%v %v %v %c[0m \n", string(b), infoColor, msg, 0x1B)
		case 3:
			ErrorColor = fmt.Sprintf("%c[%d;%d;%dm%s", 0x1B, 1, 40, 31, "")
			fmt.Printf("%v %v %v %c[0m \n", string(b), ErrorColor, msg, 0x1B)
		case 4:
			WarnColor = fmt.Sprintf("%c[%d;%d;%dm%s", 0x1B, 1, 40, 33, "")
			fmt.Printf("%v %v %v %c[0m \n", string(b), WarnColor, msg, 0x1B)
		case 7:
			DebugColor = fmt.Sprintf("%c[%d;%d;%dm%s", 0x1B, 1, 40, 34, "")
			fmt.Printf("%v %v %v %c[0m \n", string(b), DebugColor, msg, 0x1B)
		}
	}
	return nil
}
func (c *consoleWrite) destroy() {

}
func (c *consoleWrite) flush() {

}
