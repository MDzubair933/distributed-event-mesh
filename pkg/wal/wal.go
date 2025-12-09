package wal

import (
	"encoding/binary"
	"fmt"
	"os"
	"sync"
)

type Record struct {
	Offset    uint64
	Timestamp int64
	Key       []byte
	Value     []byte
}

type WriteAheadLog struct {
	mu       sync.RWMutex
	file     *os.File
	offset   uint64
	filepath string
}

func OpenWAL(path string) (*WriteAheadLog, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &WriteAheadLog{file: f, filepath: path}, nil
}

func (w *WriteAheadLog) Append(key, val []byte) (uint64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	w.offset++
	buf := make([]byte, 8+8+4+len(key)+4+len(val))
	binary.BigEndian.PutUint64(buf[0:8], w.offset)
	// Write header and payload
	_, err := w.file.Write(buf)
	return w.offset, err
}

func (w *WriteAheadLog) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.file.Close()
}
