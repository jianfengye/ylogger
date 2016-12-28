package ylogger

import (
	"io"
	"os"
)

var DefaultYLogger = NewYLogger(os.Stdout)

func Debug(class string, v ...interface{}) {
	DefaultYLogger.Debug(class, v)
}

func Info(class string, v ...interface{}) {
	DefaultYLogger.Info(class, v)
}

func Trace(class string, v ...interface{}) {
	DefaultYLogger.Trace(class, v)
}

func Warning(class string, v ...interface{}) {
	DefaultYLogger.Warning(class, v)
}

func Error(class string, v ...interface{}) {
	DefaultYLogger.Error(class, v)
}

func SetOutput(level string, w io.Writer) {
	DefaultYLogger.SetOutput(level, w)
}

func Disable(level string) {
	DefaultYLogger.Disable(level)
}

func Enable(level string) {
	DefaultYLogger.Enable(level)
}

func Close() {

}
