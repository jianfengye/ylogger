package ylogger

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

type BufioLogger struct {
	*YLogger
	writer *BufioWriter
}

type BufioWriter struct {
	file    string
	buf     *bytes.Buffer
	ticker  int
	cmd     chan int // 1 flush/ 2 close...
	cmdWait chan int
	sync.Mutex
}

func NewBufioLogger(file string, sec int) *BufioLogger {
	w := NewBufioWriter(file, sec)
	ylogger := NewYLogger(w)
	bl := new(BufioLogger)
	bl.YLogger = ylogger
	bl.writer = w
	return bl
}

func (b *BufioLogger) Close() {
	b.writer.cmd <- 2
	<-b.writer.cmdWait
}

// NewBufioWriter create new bufioWriter
func NewBufioWriter(file string, sec int) *BufioWriter {
	var buf bytes.Buffer

	bw := &BufioWriter{file: file, ticker: sec, buf: &buf}
	bw.cmd = make(chan int)
	bw.cmdWait = make(chan int)

	ticker := time.NewTicker(time.Duration(sec) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				bw.flush()
			case c := <-bw.cmd:
				switch c {
				case 1: // flush
					bw.flush()
				case 2:
					bw.flush()
					bw.cmdWait <- 1
				}
			}
		}
	}()

	return bw
}

// check file exist and write
func (f *BufioWriter) Write(b []byte) (int, error) {
	f.Lock()
	defer f.Unlock()
	return f.buf.Write(b)
}

func (bw *BufioWriter) flush() {
	bw.Lock()
	f, err := os.OpenFile(bw.file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		bw.Unlock()
	}
	bw.buf.WriteTo(f)
	bw.Unlock()
}
