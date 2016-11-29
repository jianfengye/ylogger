package ylogger

import (
	"fmt"
	"io"
	"log"
)

var LEVELS = []string{"trace", "info", "warning", "error", "debug"}

type YLogger struct {
	trace   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	debug   *log.Logger

	trace_s   bool
	info_s    bool
	warning_s bool
	err_s     bool
	debug_s   bool
}

// create new YLogger
// this out is work for all trace/info/warning/error/debug
func NewYLogger(out io.Writer) *YLogger {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	ylogger := new(YLogger)
	ylogger.trace = log.New(out, "\033[32m[TRACE]\033[0m ", flag)
	ylogger.info = log.New(out, "\033[32m[INFO]\033[0m ", flag)
	ylogger.warning = log.New(out, "\033[33m[WARNING]\033[0m ", flag)
	ylogger.err = log.New(out, "\033[31m[ERROR]\033[0m ", flag)
	ylogger.debug = log.New(out, "\033[32m[DEBUG]\033[0m ", flag)

	ylogger.trace_s = false
	ylogger.info_s = false
	ylogger.warning_s = false
	ylogger.err_s = false
	ylogger.debug_s = false
	return ylogger
}

// output:
func (this *YLogger) Debug(class string, v ...interface{}) {
	if this.debug_s {
		this.debug.Output(2, fmt.Sprint(class, " ", v))
	}
}

func (this *YLogger) Info(class string, v ...interface{}) {
	if this.info_s {
		this.info.Output(2, fmt.Sprint(class, " ", v))
	}
}

func (this *YLogger) Trace(class string, v ...interface{}) {
	if this.trace_s {
		this.trace.Output(2, fmt.Sprint(class, " ", v))
	}
}

func (this *YLogger) Warning(class string, v ...interface{}) {
	if this.warning_s {
		this.warning.Output(2, fmt.Sprint(class, " ", v))
	}
}

func (this *YLogger) Error(class string, v ...interface{}) {
	if this.err_s {
		this.err.Output(2, fmt.Sprint(class, " ", v))
	}
}

func (this *YLogger) SetOutput(level string, w io.Writer) {
	switch level {
	case "trace":
		this.trace.SetOutput(w)
	case "debug":
		this.debug.SetOutput(w)
	case "info":
		this.info.SetOutput(w)
	case "warning":
		this.warning.SetOutput(w)
	case "error":
		this.err.SetOutput(w)
	}
}

func (this *YLogger) Disable(level string) {
	switch level {
	case "trace":
		this.trace_s = false
	case "debug":
		this.debug_s = false
	case "info":
		this.info_s = false
	case "warning":
		this.warning_s = false
	case "error":
		this.err_s = false
	}
}

func (this *YLogger) Enable(level string) {
	switch level {
	case "trace":
		this.trace_s = true
	case "debug":
		this.debug_s = true
	case "info":
		this.info_s = true
	case "warning":
		this.warning_s = true
	case "error":
		this.err_s = true
	case "all":
		this.trace_s = true
		this.debug_s = true
		this.info_s = true
		this.warning_s = true
		this.err_s = true
	}
}
