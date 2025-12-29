package protocol

import "encoding/binary"

const (
	MagicByte = 0xEM
	Version1  = 1
)

type FrameHeader struct {
	Magic   byte
	Version byte
	Cmd     uint16
	Length  uint32
}

func EncodeHeader(cmd uint16, length uint32) []byte {
	buf := make([]byte, 8)
	buf[0] = 0xAA
	buf[1] = Version1
	binary.BigEndian.PutUint16(buf[2:4], cmd)
	binary.BigEndian.PutUint32(buf[4:8], length)
	return buf
}
