package account

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"godswar/pkg/types/account"
	"godswar/pkg/utility"
)

func (s service) generateUserInfo(character account.CharacterBase, kitBag account.CharacterKitBag) bytes.Buffer {
	fmt.Println(*kitBag.Equip)
	ebag := *kitBag.Equip
	equip, _, _ := s.parseKitBag(ebag, false, false)
	eblank, _, _ := s.parseKitBag(ebag, true, true)

	fmt.Println(hex.Dump(equip.Bytes()))

	kitBag1, kitBagData, _ := s.parseKitBag(kitBag.KitBag1, true, false)
	_, equipData, _ := s.parseKitBag(ebag, true, false)
	var userInfo bytes.Buffer
	var charName [32]byte
	copy(charName[:], character.Name)
	var gender uint8
	if character.Gender == "male" {
		gender = 1
	} else {
		gender = 0
	}

	pad4 := make([]byte, 4)
	x := utility.Float32Bytes(character.PosX)
	z := utility.Float32Bytes(character.PosZ)

	fmt.Println(hex.Dump(x), hex.Dump(z))

	userInfo.Write(charName[:])
	userInfo.WriteByte(gender)
	userInfo.WriteByte(uint8(character.Camp))
	userInfo.WriteByte(uint8(character.Belief))
	userInfo.WriteByte(uint8(character.Profession))
	userInfo.WriteByte(uint8(character.HairStyle))
	userInfo.WriteByte(uint8(character.FaceShape))
	userInfo.WriteByte(uint8(character.CurrentMap))
	userInfo.WriteByte(uint8(0))
	userInfo.Write([]byte{0x48, 0x14, 0x00, 0x00,})
	userInfo.Write([]byte{0x60, 0x00, 0x00, 0x00,})
	userInfo.Write(x)
	userInfo.Write(pad4)
	userInfo.Write(z)
	userInfo.Write(utility.IntByte(character.MaxHP))
	userInfo.Write(utility.IntByte(character.MaxMP))
	userInfo.Write(utility.IntByte(character.CurHP))
	userInfo.Write(utility.IntByte(character.CurMP))
	userInfo.Write(utility.IntByte(character.FighterJobExp))
	userInfo.Write(utility.IntByte(10000))
	userInfo.Write(pad4)
	userInfo.Write(pad4)
	userInfo.Write([]byte{0x3C, 0x00, 0x00, 0x00})

	equip.WriteTo(&userInfo)
	eblank.WriteTo(&userInfo)
	kitBag1.WriteTo(&userInfo)
	//padding := make([]byte, 1248)
	//userInfo.Write(padding)
	//userInfo.Write(experimental.ENTER_PART4)


	bLen := len(userInfo.Bytes())
	pvl := make([]byte, 2)
	binary.LittleEndian.PutUint16(pvl, uint16(bLen+8))

	var complete bytes.Buffer
	complete.Write(pvl)
	complete.Write([]byte{0x23, 0x27, 0xD2, 0xE4, 0xEB, 0x0B})
	userInfo.WriteTo(&complete)

	cb := complete.Bytes()
	s.conn.Bag1 = &kitBagData
	s.conn.Equip = &equipData
	s.conn.CharacterBytes = &cb
	fmt.Println(len(cb))
	return complete
}

/*
	Generate user info for user preview
 */
func (s service) generateUserPreview (character account.CharacterBase, equipIDs bytes.Buffer) bytes.Buffer {

	var charPreview bytes.Buffer
	var charName [32]byte
	copy(charName[:], character.Name)

	var gender uint8
	if character.Gender == "male" {
		gender = 1
	} else {
		gender = 0
	}

	pad := make([]byte, 47)
	equipIDs.Write(pad)

	charPreview.Write(charName[:])
	charPreview.WriteByte(uint8(character.Camp))
	charPreview.WriteByte(uint8(character.Profession))
	charPreview.WriteByte(uint8(character.FighterJobLv))
	charPreview.WriteByte(gender)
	charPreview.WriteByte(uint8(character.HairStyle))
	charPreview.WriteByte(uint8(character.FaceShape))
	charPreview.WriteByte(uint8(0))
	equipIDs.WriteTo(&charPreview)

	previewLen := len(charPreview.Bytes())
	pvl := make([]byte, 2)
	binary.LittleEndian.PutUint16(pvl, uint16(previewLen+5))
	var charPreviewResponse bytes.Buffer
	charPreviewResponse.Write(pvl)
	charPreviewResponse.Write([]byte{0x12, 0x27, 0x01})
	charPreview.WriteTo(&charPreviewResponse)

	return charPreviewResponse
}