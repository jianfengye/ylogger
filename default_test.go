package ylogger

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func BenchmarkDevNull(b *testing.B) {
	var str string = "ykafka start manager"
	blogger := NewYLogger(ioutil.Discard)
	blogger.Enable("all")
	for i := 0; i < b.N; i++ {
		blogger.Info("bench blogger", str)
	}
}

func TestDefault(t *testing.T) {
	Enable("all")
	var buf bytes.Buffer
	foo := bufio.NewWriter(&buf)
	SetOutput("debug", foo)

	Debug("test logger", "test", "test2")

	out := buf.String()
	fmt.Println(out)
}
