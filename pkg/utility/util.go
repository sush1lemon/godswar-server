package utility

import (
	"encoding/binary"
	"time"
)

func IntByte(v int) []byte  {
	x := make([]byte, 4)
	binary.LittleEndian.PutUint16(x, uint16(v))
	return x
}

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}