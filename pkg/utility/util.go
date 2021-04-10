package utility

import "encoding/binary"

func IntByte(v int) []byte  {
	x := make([]byte, 4)
	binary.LittleEndian.PutUint16(x, uint16(v))
	return x
}
