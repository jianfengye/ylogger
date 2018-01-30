package ylogger

import "testing"

func BenchmarkChan(b *testing.B) {
	var str string = "ykafka start manager"
	blogger := NewChanLogger("/tmp/degrade/chan.log", 2, 3000)
	blogger.Enable("debug")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		blogger.Debug("kafka start", str, "test2")
	}
	blogger.Close()
	b.StopTimer()
}
