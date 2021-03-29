package utility

import (
	"bytes"
	"encoding/binary"
)

var cursor = 0

type packetreader struct {
	buffer []byte
}

func (p packetreader) UInt32() uint32 {
	tmp := make([]byte, len(p.buffer) - 2)
	copy(tmp, p.buffer[cursor + 2:])


	x := bytes.NewBuffer(tmp)
	x.ReadBytes(2)

	binary.Read(bytes.NewReader(tmp), binary.LittleEndian, x)

	return binary.LittleEndian.Uint32(tmp)
}

func (p packetreader) Skip(len int) {
	cursor += len
}

func (p packetreader) Read(len int) string {
	var to int
	if cursor == 0 {
		to = (len + cursor) - 1
	} else {
		to = len + cursor
	}

	buff := make([]byte, len)
	copy(buff, p.buffer[cursor:to])
	cursor += len
	return string(buff)
}

func (p packetreader) Len() int  {
	return len(p.buffer)
}

type PacketReader interface {
	Read(len int) string
	Skip(len int)
	UInt32() uint32
	Len() int
}

func NewPacketReader(buffer []byte) PacketReader {
	cursor = 0
	return &packetreader{buffer}
}
