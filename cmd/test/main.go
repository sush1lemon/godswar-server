package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"godswar/pkg/packets"
	"math"
)

type Test struct {
	Username  [32]byte
	Password  [32]byte
	Unk1      [4]byte
	ClientMac [32]byte
	Unk2      [40]byte
}

func main()  {

	fmt.Println(len(packets.ENTER_PART1), len(packets.ENTER_PART4), len(packets.CHAMPTEST))
	bs := []byte{0x58, 0x06, 0x23, 0x27,}
	wnt := make([]byte, 4)
	bc := binary.LittleEndian.Uint16(bs)
	fmt.Println(bc)
	fmt.Println(math.Float32frombits(uint32(bc)))
	binary.LittleEndian.PutUint16(wnt, 2030)

	fmt.Println(hex.Dump(bs))
	fmt.Println(hex.Dump(wnt))
	//fmt.Println(hex.Dump([]byte("abcdefghijklmnopqrstuvwxyz123456")))

	//packets.Test()



	//var test Test
	//buf := bytes.NewReader(h.decoded.DecodedBuffer[:])
	//binary.Read(buf, binary.LittleEndian, &test)
	//fmt.Println(string(test.Password[:]))
	//fmt.Println(utility.Parser(string(test.Username[:])))

	//0x01, // POLIS 00 - sparta 01 -athens
	//0x03, // CLASS 00 - wr 01 - ch - 02 -pr -03 -mgc
	//0x8C, // LEVEL
	//0x01, // GENDER
	//0x35, // HAIR
	//0x00, // FACE
}