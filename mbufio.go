package ylogger

import (
	"os"
)

func NewMBufioYLogger(base string) (*YLogger, error) {
	ylogger := NewYLogger(os.Stdout)

	for _, level := range LEVELS {
		file := base + "." + level
		w, err := NewBufioWriter(file)
		if err != nil {
			return nil, err
		}
		ylogger.SetOutput(level, w)
	}

	return ylogger, nil
}
