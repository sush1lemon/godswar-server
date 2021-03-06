package account

import (
	"github.com/upper/db/v4"
	"godswar/pkg/decode"
	"godswar/pkg/networking"
)

type Service interface {
	Login(packet *decode.Decode)
	LoginFail(reason uint16)
	GetAccountCharacters(packet *decode.Decode)
	CreateAccountCharacter(packet *decode.Decode)
	DeleteAccountCharacter(packet decode.Decode)
}

type service struct {
	db   db.Session
	conn *networking.Connection
}

func NewAccountService(db db.Session, conn *networking.Connection) Service {
	return &service{db, conn}
}

type Account struct {
	Username string
	Password string
}

type Server struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	IP          string `db:"ip_address"`
	Identifier  string `db:"identifier"`
	ServerLimit int    `db:"server_limit"`
}
