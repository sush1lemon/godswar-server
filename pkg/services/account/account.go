package account

import (
	"database/sql"
	"godswar/pkg/decode"
	"godswar/pkg/networking"
)

type Service interface {
	Login(packet *decode.Decode)
	LoginFail(reason uint16)
}

type service struct {
	db   *sql.DB
	conn *networking.Connection
}

type Account struct {
	Username string
	Password string
}

type AuthRequest struct {
	Username  [32]byte
	Password  [32]byte
	Unk1      [4]byte
	ClientMac [32]byte
}

type GSAuthRequest struct {
	Username [32]byte
	ServerID byte
	Unk      [3]byte
}

type GetCharacterRequest struct {
	Username [32]byte
	Unk      [24]byte
}
