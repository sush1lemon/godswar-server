package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"godswar/pkg/packets"
	"time"
)

type Test struct {
	Username  [32]byte
	Password  [32]byte
	Unk1      [4]byte
	ClientMac [32]byte
	Unk2      [40]byte
}
func getSecondOfDay(t time.Time) int {
	return 60*60*t.Hour() + 60*t.Minute() + t.Second()
}
func main() {

	t := time.Now()
	ts := getSecondOfDay(t)

	fmt.Println([]byte("[4000,,,,,,0,10,1,1,0]#[4030,,,,,,0,10,1,1,0]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#[]#"))


	fmt.Println(len(packets.ENTER_PART1), len(packets.ENTER_PART4), len(packets.CHAMPTEST))
	bs := []byte{0x9C, 0xBE, 0x64, 0x60,}
	wnt := make([]byte, 4)
	bc := binary.LittleEndian.Uint16(bs)
	fmt.Println(bc)
	binary.LittleEndian.PutUint16(wnt, uint16(ts))

	//fmt.Println(hex.Dump(bs))
	fmt.Println(hex.Dump(wnt))
	fmt.Println(ts)
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
