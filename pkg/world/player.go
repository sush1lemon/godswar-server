package world

type Player struct {
	Name    [32]byte
	Gender  byte
	Faction byte
	Faith   byte
	Class   byte
	Hair    byte
	Unk1    byte
	Unk2    byte
	Unk3    byte
	UnkA1   [4]byte
	UnkA2   [4]byte
	CoorX   [4]byte
	CoorZ   [4]byte
	CoorY   [4]byte
	MaxHP   [4]byte
	MaxMP   [4]byte
	CurHP   [4]byte
	CurMP   [4]byte
	ITEMS   []EquipItem
}

type EquipItem struct {
	ItemID [4]byte
	Attr1  [4]byte
	Attr2  [4]byte
	Attr3  [4]byte
	Attr4  [4]byte
	Attr5  [4]byte
}

func NewPlayer() Player {
	return Player{}
}
