package account

import (
	"bytes"
	"database/sql"
	"encoding/binary"
	"fmt"
	"godswar/pkg/decode"
	"godswar/pkg/networking"
	"godswar/pkg/packets"
	"godswar/pkg/utility"
)

func NewAccountService(db *sql.DB, conn *networking.Connection) Service {
	return &service{db, conn}
}

func (s service) LoginFail(reason uint16) {
	bar := new(bytes.Buffer)
	bar.Write([]byte{6})
	bar.Write([]byte{0})
	err := binary.Write(bar, binary.LittleEndian, reason)
	if err != nil {
		fmt.Println(err)
		s.conn.Disconnect()
	}
	bar.Write([]byte{0x00, 0x00, 0xf0})
	s.conn.Send(bar.Bytes())
}

func (s service) Login(packet *decode.Decode) {
	var request AuthRequest
	var account Account
	reader := bytes.NewReader(packet.DecodedBuffer[:])
	err := binary.Read(reader, binary.LittleEndian, &request)
	if err != nil {
		fmt.Println(err)
		s.conn.Disconnect()
	}
	username := utility.Parser(string(request.Username[:]))
	password := string(request.Password[:])

	query, err := s.db.Prepare("SELECT username, password from accounts WHERE username = ? AND password= ?")
	if err != nil {
		s.conn.Disconnect()
	}
	defer query.Close()

	row := query.QueryRow(&username, &password)
	switch err := row.Scan(&account.Username, &account.Password); err {
	case nil:
		s.conn.Send(packets.SERVER_LIST)
		break
	default:
		s.LoginFail(3)
	}
}

func (s service) GetAccountCharacters(packet decode.Decode) {

}