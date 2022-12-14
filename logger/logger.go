package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// RFC5424 log message levels.
const (
	LevelEmergency     = iota //紧急的
	LevelAlert                //警报
	LevelCritical             //重要的
	LevelError                //错误
	LevelWarning              //警告
	LevelNotice               //提示
	LevelInformational        //信息
	LevelDebug                //调试
)

const (
	LevelTrace = LevelDebug
	LevelWarn  = LevelWarning
	LevelInfo  = LevelInformational
)

const (
	AdapterConsole   = "console"
	AdapterFile      = "file"
	AdapterMultiFile = "multifile"
)

func register(adapterName string, log newLoggerFunc) {
	if nil == log {
		panic("logs: 注册类型函数为空")
	}
	if _, dup := adapters[adapterName]; dup {
		fmt.Println(fmt.Sprintf("logs：日志类型：%v,已经注册", adapterName))
		//fmt.Fprintln(os.Stderr, "logs：日志类型：%s,已经注册", adapterName)
		return
	}
	adapters[adapterName] = log
}

type newLoggerFunc func() logger

var adapters = make(map[string]newLoggerFunc)
var levelPrefix = [LevelDebug + 1]string{"[M] ", "[A] ", "[C] ", "[E] ", "[W] ", "[N] ", "[I] ", "[D] "}
var levelName = [LevelDebug + 1]string{"Emergency",
	"Alert",
	"Critical",
	"Error",
	"Warning",
	"Notice",
	"Informational",
	"Debug"}

type logger interface {
	init() error
	writeMsg(when time.Time, msg string, level int) error
	destroy()
	flush()
}

type loggers struct {
	lock         sync.Mutex
	consolePluse bool //true console暂停日志滚动
	//level int
	//init                bool
	enableFuncCallDepth bool
	loggerFuncCallDepth int
	asynchronous        bool
	msgChanLen          int64
	msgChan             chan *logMsg
	signalChan          chan string
	outputs             []*nameLogger
}

const defaultAsyncMsgLen = 1e3

type nameLogger struct {
	logger
	level int
	name  string
}

type logMsg struct {
	level               int
	msg                 string
	when                time.Time
	customFuncCallDepth int //自定义函数路径深度
}

func newLogger(channelLens ...int64) *loggers {
	l := new(loggers)
	//l.level = LevelDebug
	l.consolePluse = false
	l.enableFuncCallDepth = true
	l.loggerFuncCallDepth = 3
	l.msgChanLen = append(channelLens, 0)[0]
	if l.msgChanLen <= 0 {
		l.msgChanLen = defaultAsyncMsgLen
	}
	l.signalChan = make(chan string, 1)

	//默认注册CONSOLE类日志
	//register(AdapterConsole, newConsole)
	//l.setLogger(AdapterConsole)
	return l
}

func (l *loggers) async(msgLen int64) *loggers {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.asynchronous {
		return l
	}

	l.asynchronous = true
	if msgLen > 0 {
		l.msgChanLen = msgLen
	}
	l.msgChan = make(chan *logMsg, l.msgChanLen)
	go l.startLogger()
	return l
}

func (l *loggers) setLogger(adapterName string) error {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, log := range l.outputs {
		if log.name == adapterName {
			err := fmt.Errorf("logs:你已经设置了这个日志名称， %v", adapterName)
			fmt.Fprintln(os.Stderr, err)
			return err
		}
	}

	log, ok := adapters[adapterName]
	if !ok {
		err := fmt.Errorf("logs:设置日志错误，该类型未注册，%v", adapterName)
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	lg := log()
	err := lg.init()
	if err != nil {
		fmt.Fprintln(os.Stderr, "logs.loggers.setLogger:", err.Error())
		return err
	}
	l.outputs = append(l.outputs, &nameLogger{name: adapterName, level: LevelDebug, logger: lg})
	return nil
}

func (l *loggers) delLogger(adapterName string) {
	l.lock.Lock()
	defer l.lock.Unlock()
	outputs := []*nameLogger{}
	for _, lg := range l.outputs {
		if lg.name == adapterName {
			lg.destroy()
		} else {
			outputs = append(outputs, lg)
		}
	}
	l.outputs = outputs
}

func (l *loggers) setLevel(adapterName string, level int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	for _, lg := range l.outputs {
		if lg.name == adapterName {
			lg.level = level
		}
	}
}

func (l *loggers) setConsolePluse(pluse bool) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.consolePluse = pluse
}

func (l *loggers) startLogger() {
	gameOver := false
	for {
		select {
		case bm := <-l.msgChan:
			l.writeTologgers(bm.when, bm.msg, bm.level)

		case sg := <-l.signalChan:

			if sg == "close" {
				for _, lg := range l.outputs {
					lg.destroy()
				}
				l.outputs = nil
				gameOver = true
			}
		}
		if gameOver {
			break
		}
	}
}

func (l *loggers) writeTologgers(when time.Time, msg string, level int) {
	for _, lg := range l.outputs {
		if level > lg.level {
			continue
		}
		if lg.name == AdapterConsole &&
			l.consolePluse {
			continue
		}
		err := lg.writeMsg(when, msg, level)
		if err != nil {
			fmt.Fprintf(os.Stderr, "不能输出到adapter: %v，error：%v\n", lg.name, err)
		}
	}
}

func (l *loggers) writeMsg(customFuncCallDepth int, level int, msg string, v ...interface{}) {
	when := time.Now()
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	if level != LevelInfo && level != LevelNotice && level != LevelWarn {
		if l.enableFuncCallDepth {
			funcCallDepth := l.loggerFuncCallDepth
			if customFuncCallDepth > 0 {
				funcCallDepth = customFuncCallDepth
			}
			pc, file, line, ok := runtime.Caller(funcCallDepth)
			if !ok {
				file = "???"
				line = 0
			}
			var funcPath string = ""
			if f := runtime.FuncForPC(pc); f != nil {
				funcPath = f.Name()
			}
			_, filename := path.Split(file)
			msg = msg + "	[" + funcPath + " " + filename + ":" + strconv.FormatInt(int64(line), 10) + "]"
		}
	}

	msg = levelPrefix[level] + msg

	if GetRequestId() != nil {
		msg = fmt.Sprintf("[requestId: %v] \t", GetRequestId()) + msg
	}

	if l.asynchronous {
		lm := new(logMsg)
		lm.level = level
		lm.msg = msg
		lm.when = when
		l.msgChan <- lm
	} else {
		l.writeTologgers(when, msg, level)
	}
}
func (l *loggers) Emergency(customFuncCallDepth int, format string, v ...interface{}) {
	l.writeMsg(customFuncCallDepth, LevelEmergency, format, v...)
}
func (l *loggers) Alert(customFuncCallDepth int, format string, v ...interface{}) {
	l.writeMsg(customFuncCallDepth, LevelAlert, format, v...)
}
func (l *loggers) Critical(customFuncCallDepth int, format string, v ...interface{}) {
	l.writeMsg(customFuncCallDepth, LevelCritical, format, v...)
}
func (l *loggers) Error(customFuncCallDepth int, format string, v ...interface{}) {
	l.writeMsg(customFuncCallDepth, LevelError, format, v...)
}
func (l *loggers) Warning(customFuncCallDepth int, format string, v ...interface{}) {
	l.writeMsg(customFuncCallDepth, LevelWarning, format, v...)
}
func (l *loggers) Notice(customFuncCallDepth int, format string, v ...interface{}) {
	l.writeMsg(customFuncCallDepth, LevelNotice, format, v...)
}
func (l *loggers) Informational(customFuncCallDepth int, format string, v ...interface{}) {
	l.writeMsg(customFuncCallDepth, LevelInformational, format, v...)
}
func (l *loggers) Debug(customFuncCallDepth int, format string, v ...interface{}) {
	l.writeMsg(customFuncCallDepth, LevelDebug, format, v...)
}
