package benchmarks

import (
	"testing"
	"github.com/MDzubair933/distributed-event-mesh/pkg/wal"
)

func BenchmarkWALAppend(b *testing.B) {
	log, _ := wal.OpenWAL("/tmp/bench.wal")
	defer log.Close()
	data := []byte("benchmark-payload-bytes")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		log.Append([]byte("key"), data)
	}
}
