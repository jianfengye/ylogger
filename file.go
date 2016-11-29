package ylogger

import (
	"os"
)

type FileWriter struct {
	file string
	b    *os.File
}

func NewFileYLogger(file string) (*YLogger, error) {
	w, err := NewFileWriter(file)
	if err != nil {
		return nil, err
	}

	return NewYLogger(w), nil
}

func NewFileWriter(file string) (*FileWriter, error) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &FileWriter{file: file, b: f}, nil
}

// check file exist and write
func (f *FileWriter) Write(b []byte) (int, error) {
	// TODO: 这里是不是可以考虑使用watcher修改
	/*
	   _, err := f.b.Stat()
	   if err != nil {
	       f.b, err = os.OpenFile(f.file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	       if err != nil {
	           panic("reopen file error: " + err.Error())
	       }
	   }
	*/

	return f.b.Write(b)
}
