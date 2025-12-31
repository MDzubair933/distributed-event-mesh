package tests

import (
	"testing"
	"github.com/MDzubair933/distributed-event-mesh/pkg/protocol"
)

func FuzzProtocolFraming(f *testing.F) {
	f.Add(uint16(10), uint32(1024))
	f.Fuzz(func(t *testing.T, cmd uint16, length uint32) {
		hdr := protocol.EncodeHeader(cmd, length)
		if len(hdr) != 8 {
			t.Errorf("Unexpected header length %d", len(hdr))
		}
	})
}
