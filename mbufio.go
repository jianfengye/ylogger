package ylogger

import "os"

type MBufioLogger struct {
	*YLogger

	wTrace   *BufioWriter
	wDebug   *BufioWriter
	wInfo    *BufioWriter
	wWarning *BufioWriter
	wError   *BufioWriter
}

func NewMBufioYLogger(base string, sec int) *MBufioLogger {
	ylogger := NewYLogger(os.Stdout)

	mb := new(MBufioLogger)
	mb.YLogger = ylogger

	mb.wDebug = NewBufioWriter(base+"."+"debug", sec)
	mb.YLogger.SetOutput("debug", mb.wDebug)

	mb.wTrace = NewBufioWriter(base+"."+"trace", sec)
	mb.YLogger.SetOutput("trace", mb.wTrace)

	mb.wInfo = NewBufioWriter(base+"."+"info", sec)
	mb.YLogger.SetOutput("info", mb.wInfo)

	mb.wWarning = NewBufioWriter(base+"."+"warning", sec)
	mb.YLogger.SetOutput("warning", mb.wWarning)

	mb.wError = NewBufioWriter(base+"."+"error", sec)
	mb.YLogger.SetOutput("error", mb.wError)

	return mb
}

func (mb *MBufioLogger) Close() {
	mb.wDebug.cmd <- 1
	mb.wTrace.cmd <- 1

	mb.wInfo.cmd <- 1
	mb.wWarning.cmd <- 1
	mb.wError.cmd <- 1
}
