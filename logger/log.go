package logger

import (
	"fmt"
	"os"
)

var myLogger *loggers = newLogger()

func RegAdapter(adapterName string) {
	switch adapterName {
	case AdapterConsole:
		register(AdapterConsole, newConsole)
		myLogger.setLogger(AdapterConsole)
		break
	case AdapterFile:
		register(AdapterFile, newFile)
		myLogger.setLogger(AdapterFile)
		break
	case AdapterMultiFile:
		register(AdapterMultiFile, newMultiFile)
		myLogger.setLogger(AdapterMultiFile)
		break
	default:
		fmt.Fprintln(os.Stderr, "logs:注册的名称未实现:%v", adapterName)
		break
	}
}

func UnRegAdapter(adapterName string) {
	myLogger.delLogger(adapterName)
	delete(adapters, adapterName)
}

func SetLoggerLevel(adapterName string, level int) {
	if level < LevelEmergency {
		level = LevelEmergency
	}
	if level > LevelDebug {
		level = LevelDebug
	}
	myLogger.setLevel(adapterName, level)
}

func Async(msgChanLen int64) {
	myLogger.async(msgChanLen)
}

func Emergency(f interface{}, v ...interface{}) {
	myLogger.Emergency(0, formatLog(f, v...))
}

func Alert(f interface{}, v ...interface{}) {
	myLogger.Alert(0, formatLog(f, v...))
}

func Critical(f interface{}, v ...interface{}) {
	myLogger.Critical(0, formatLog(f, v...))
}
func Error(f interface{}, v ...interface{}) {
	myLogger.Error(0, formatLog(f, v...))
}
func Warning(f interface{}, v ...interface{}) {
	myLogger.Warning(0, formatLog(f, v...))
}
func Notice(f interface{}, v ...interface{}) {
	myLogger.Notice(0, formatLog(f, v...))
}
func Informational(f interface{}, v ...interface{}) {
	myLogger.Informational(0, formatLog(f, v...))
}

func Debug(f interface{}, v ...interface{}) {
	myLogger.Debug(0, formatLog(f, v...))
}

func Warn(f interface{}, v ...interface{}) {
	myLogger.Warning(0, formatLog(f, v...))
}
func Info(f interface{}, v ...interface{}) {
	myLogger.Informational(0, formatLog(f, v...))
}
func Trace(f interface{}, v ...interface{}) {
	myLogger.Debug(0, formatLog(f, v...))
}

//*****************************************************
func Emergency_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Emergency(funcCallDepth, formatLog(f, v...))
}

func Alert_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Alert(funcCallDepth, formatLog(f, v...))
}

func Critical_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Critical(funcCallDepth, formatLog(f, v...))
}
func Error_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Error(funcCallDepth, formatLog(f, v...))
}
func Warning_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Warning(funcCallDepth, formatLog(f, v...))
}
func Notice_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Notice(funcCallDepth, formatLog(f, v...))
}
func Informational_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Informational(funcCallDepth, formatLog(f, v...))
}
func Debug_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Debug(funcCallDepth, formatLog(f, v...))
}
func Warn_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Warning(funcCallDepth, formatLog(f, v...))
}
func Info_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Informational(funcCallDepth, formatLog(f, v...))
}
func Trace_CustomFuncCallDepth(funcCallDepth int, f interface{}, v ...interface{}) {
	myLogger.Debug(funcCallDepth, formatLog(f, v...))
}

//*****************************************************

func ConsoleStopShow() {
	myLogger.setConsolePluse(true)
}

func ConsoleAgainShow() {
	myLogger.setConsolePluse(false)
}
