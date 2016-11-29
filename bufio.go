package ylogger

import (
	"bufio"
	"os"
	"sync"
	"time"
)

type BufioWriter struct {
	file   string
	b      *bufio.Writer
	ticker *time.Ticker
	sync.Mutex
}

func NewBufioYLogger(file string, ticker int64) (*YLogger, error) {
	w, err := NewBufioWriter(file)
	if err != nil {
		return nil, err
	}

	w.ticker = time.NewTicker(time.Duration(ticker) * time.Second)

	go func() {
		for {
			select {
			case <-w.ticker.C:
				w.b.Flush()
			}
		}
	}()

	return NewYLogger(w), nil
}

// NewBufioWriter 新创建一个bufioWriter
func NewBufioWriter(file string) (*BufioWriter, error) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	bf := bufio.NewWriter(f)
	bufWriter := &BufioWriter{file: file, b: bf}

	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				bufWriter.Lock()
				if bufWriter.b == nil {
					f2, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
					if err != nil {
						return
					}
					bf2 := bufio.NewWriter(f2)
					bufWriter.b = bf2
				}
				bufWriter.b.Flush()
				bufWriter.Unlock()
			}
		}
	}()

	return &BufioWriter{file: file, b: bf}, nil
}

// check file exist and write
func (f *BufioWriter) Write(b []byte) (int, error) {
	f.Lock()
	defer f.Unlock()
	return f.b.Write(b)
}
