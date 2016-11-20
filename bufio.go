package ylogger

import(
    "bufio"
    "os"
    "time"
)

type BufioWriter struct {
    file string
    b *bufio.Writer
    ticker *time.Ticker
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
            case <- w.ticker.C:
                    w.b.Flush()
            }
        }
    }()

    return NewYLogger(w), nil
}

func NewBufioWriter(file string) (*BufioWriter, error) {
    f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

    bf := bufio.NewWriter(f)

    return &BufioWriter{file: file, b: bf}, nil
}

// check file exist and write
func (f *BufioWriter) Write(b []byte) (int, error) {
    //TODO: 检测file是否有改变
    return f.b.Write(b)
}
