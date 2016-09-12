package ylogger

import (
	"fmt"
	"io"
	"log"
)

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

const (
	FgBlack = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// create new YLogger
// this out is work for all trace/info/warning/error/debug
func NewYLogger(out io.Writer) *YLogger {
	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile
	ylogger := new(YLogger)
	ylogger.trace = log.New(out, fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgYellow, "[TRACE] "), flag)
	ylogger.info = log.New(out, fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgGreen, "[INFO] "), flag)
	ylogger.warning = log.New(out, fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgBlue, "[WARNING] "), flag)
	ylogger.err = log.New(out, fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgRed, "[ERROR] "), flag)
	ylogger.debug = log.New(out, fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgMagenta, "[DEBUG] "), flag)

	ylogger.trace_s = true
	ylogger.info_s = true
	ylogger.warning_s = true
	ylogger.err_s = true
	ylogger.debug_s = true
	return ylogger
}

// output:
func (this *YLogger) Debug(class string, v ...interface{}) {
	if this.debug_s {
		this.debug.Output(2, fmt.Sprint(fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgGreen, class), " ", v))
	}
}

func (this *YLogger) Info(class string, v ...interface{}) {
	if this.info_s {
		this.info.Output(2, fmt.Sprint(fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgGreen, class), " ", v))
	}
}

func (this *YLogger) Trace(class string, v ...interface{}) {
	if this.trace_s {
		this.trace.Output(2, fmt.Sprint(fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgGreen, class), " ", v))
	}
}

func (this *YLogger) Warning(class string, v ...interface{}) {
	if this.warning_s {
		this.warning.Output(2, fmt.Sprint(fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgGreen, class), " ", v))
	}
}

func (this *YLogger) Error(class string, v ...interface{}) {
	if this.err_s {
		this.err.Output(2, fmt.Sprint(fmt.Sprintf("\x1b[%dm%s\x1b[0m", FgGreen, class), " ", v))
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
	}
}
