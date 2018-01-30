package ylogger

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"time"
)

// use chan for log
type CLogger struct {
	file    string
	in      chan string
	ins     chan string
	inb     chan []byte
	cmd     chan int // command channel : 1 flush/ 2 close...
	cmdWait chan int
	buf     *bytes.Buffer
	mu      sync.Mutex

	trace_s   bool
	info_s    bool
	warning_s bool
	error_s   bool
	debug_s   bool
}

// NewCLogger create a new Logger
func NewChanLogger(file string, sec int, chanbuf int) *CLogger {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil
	}

	// instance
	c := &CLogger{file: file}
	c.buf = &bytes.Buffer{}
	c.in = make(chan string, chanbuf)
	c.cmd = make(chan int)
	c.cmdWait = make(chan int)
	c.Disable("all")

	go func() {
		ticker := time.NewTicker(time.Duration(sec) * time.Second)
		for {
			select {
			case cmd := <-c.cmd:
				switch cmd {
				case 1: // flush
					c.buf.WriteTo(f)
					c.buf.Truncate(0)
				case 2: // close
					c.buf.WriteTo(f)
					c.buf.Truncate(0)
					c.cmdWait <- 1
				}
			case msg := <-c.in: // msg come
				c.buf.WriteString(msg)
				c.buf.WriteByte('\n')

			case <-ticker.C: // need flush
				// check file exist
				_, err := f.Stat()
				if err != nil {
					f, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
					if err != nil {
						fmt.Println(err)
						continue
					}
				}

				// flush
				c.buf.WriteTo(f)
				c.buf.Truncate(0)

			}
		}
	}()
	return c
}

func (this *CLogger) Disable(level string) {
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
		this.error_s = false
	case "all":
		this.trace_s = false
		this.debug_s = false
		this.info_s = false
		this.warning_s = false
		this.error_s = false
	}
}

func (this *CLogger) Enable(level string) {
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
		this.error_s = true
	case "all":
		this.trace_s = true
		this.debug_s = true
		this.info_s = true
		this.warning_s = true
		this.error_s = true
	}
}

func (this *CLogger) Close() {
	this.cmd <- 2
	<-this.cmdWait
}

func (this *CLogger) Trace(class string, v ...interface{}) {
	if this.trace_s {
		this.in <- this.output("[Trace]", fmt.Sprint(class, ' ', v))
	}
}

func (this *CLogger) Debug(class string, v ...interface{}) {
	if this.debug_s {
		this.in <- this.output("[DEBUG]", fmt.Sprint(class, ' ', v))
	}
}

func (this *CLogger) Info(class string, v ...interface{}) {
	if this.info_s {
		this.in <- this.output("[INFO]", fmt.Sprint(class, ' ', v))
	}
}

func (this *CLogger) Warning(class string, v ...interface{}) {
	if this.warning_s {
		this.in <- this.output("[WARNING]", fmt.Sprint(class, ' ', v))
	}
}

func (this *CLogger) Error(class string, v ...interface{}) {
	if this.error_s {
		this.in <- this.output("[ERROR]", fmt.Sprint(class, ' ', v))
	}
}

func (this *CLogger) output(prefix string, s string) string {
	var b []byte

	b = append(b, []byte(prefix)...)
	b = append(b, ' ')

	// append time
	t := time.Now()

	year, month, day := t.Date()
	itoa(&b, year, 4)
	b = append(b, '/')
	itoa(&b, int(month), 2)
	b = append(b, '/')
	itoa(&b, day, 2)
	b = append(b, ' ')

	hour, min, sec := t.Clock()
	itoa(&b, hour, 2)
	b = append(b, ':')
	itoa(&b, min, 2)
	b = append(b, ':')
	itoa(&b, sec, 2)
	b = append(b, '.')
	itoa(&b, t.Nanosecond()/1e3, 6)

	b = append(b, ' ')

	b = append(b, []byte(s)...)

	return string(b)
}

func itoa(buf *[]byte, i int, wid int) {
	var u uint = uint(i)
	if u == 0 && wid <= 1 {
		*buf = append(*buf, '0')
		return
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; u > 0 || wid > 0; u /= 10 {
		bp--
		wid--
		b[bp] = byte(u%10) + '0'
	}
	*buf = append(*buf, b[bp:]...)
}
