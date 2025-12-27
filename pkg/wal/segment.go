package wal

import (
	"fmt"
	"os"
)

type Segment struct {
	BaseOffset uint64
	File       *os.File
	Size       int64
	MaxBytes   int64
}

func NewSegment(dir string, baseOffset uint64, maxBytes int64) (*Segment, error) {
	filename := fmt.Sprintf("%s/%020d.log", dir, baseOffset)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	fi, _ := f.Stat()
	return &Segment{
		BaseOffset: baseOffset,
		File:       f,
		Size:       fi.Size(),
		MaxBytes:   maxBytes,
	}, nil
}
