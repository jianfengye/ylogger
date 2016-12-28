package ylogger

import "testing"

func BenchmarkBufio(b *testing.B) {
	var str string = "ykafka start manager"
	blogger := NewBufioLogger("/tmp/degrade/bufio.log", 2)
	blogger.Enable("all")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogger.Info("bench blogger", str, "test2")
	}
	blogger.Close()
	b.StopTimer()
}
