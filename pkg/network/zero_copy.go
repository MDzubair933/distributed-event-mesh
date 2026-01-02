package network

import (
	"net"
	"os"
	"syscall"
)

func SendfileToConn(conn net.Conn, file *os.File, offset int64, count int64) (int64, error) {
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
		return 0, nil
	}
	raw, err := tcpConn.SyscallConn()
	if err != nil {
		return 0, err
	}
	var written int
	raw.Control(func(fd uintptr) {
		written, _ = syscall.Sendfile(int(fd), int(file.Fd()), &offset, int(count))
	})
	return int64(written), nil
}
