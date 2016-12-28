package ylogger

import (
	"testing"
	"time"
)

func BenchmarkMBufio(b *testing.B) {
	var str string = "ykafka start manager"
	blogger := NewMBufioYLogger("/tmp/degrade/mbufio.log", 10)
	blogger.Enable("all")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogger.Info("bench blogger", str)
	}
	blogger.Close()
	b.StopTimer()
}

func TestMBufio(t *testing.T) {
	var str string = "ykafka start manager"
	blogger := NewMBufioYLogger("/tmp/degrade/mbufio.log", 1)
	blogger.Enable("all")
	blogger.Info("bench blogger", str)
	blogger.Trace("bench blogger", str)
	time.Sleep(2 * time.Second)
}
