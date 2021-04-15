package requests

type GSAuthRequest struct {
	Username [32]byte
	ServerID byte
	Unk      [3]byte
}

type GetCharacterRequest struct {
	Username         [32]byte
	ServerIdentifier [25]byte
}

type CreateCharacterRequest struct {
	CharName [32]byte
	Gender   byte
	Camp     byte
	Class    byte
	Unk      byte
	Hair     byte
	Face     byte
	Unk1     [32]byte
	Faith    byte
	Unk2     [2]byte
}

type AuthRequest struct {
	Username  [32]byte
	Password  [32]byte
	Unk1      [4]byte
	ClientMac [32]byte
}

type DeleteCharRequest struct {
	Username [32]byte
	CharName [32]byte
}
