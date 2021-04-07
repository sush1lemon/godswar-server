package account

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/upper/db/v4"
	"godswar/pkg/decode"
	"godswar/pkg/logger"
	"godswar/pkg/networking"
	"godswar/pkg/packets"
	"godswar/pkg/utility"
)

func NewAccountService(db db.Session, conn *networking.Connection) Service {
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
	reader := bytes.NewReader(packet.DecodedBuffer[:])
	err := binary.Read(reader, binary.LittleEndian, &request)
	if err != nil {
		fmt.Println(err)
		s.conn.Disconnect()
	}
	username := utility.Parser(string(request.Username[:]))
	password := string(request.Password[:])

	col := s.db.Collection("accounts")
	q := col.Find(db.Cond{"username": username, "password": password})
	total, err := q.Count()
	if err != nil {
		logger.BasicLog(err)
		s.conn.Disconnect()
	}

	if total > 0 {
		s.conn.Send(packets.SERVER_LIST)
		return
	} else {
		s.LoginFail(3)
	}
}

func (s service) GetAccountCharacters(packet *decode.Decode) {
	var request GetCharacterRequest
	var character CharacterBase
	var server Server
	reader := bytes.NewReader(packet.Buffer[4:])
	err := binary.Read(reader, binary.LittleEndian, &request)
	if err != nil {
		logger.BasicLog(err)
		s.conn.Disconnect()
	}

	s.conn.ServerID = &server.ID
	username := utility.RemoveBlank(string(request.Username[:]))
	identifier := string(request.ServerIdentifier[:])

	err = s.db.SQL().Select("*").From("server").
		Where("identifier", identifier).One(&server)

	if err != nil {
		logger.BasicLog("Server not found", err)
		s.conn.Disconnect()
	}

	query := s.db.SQL().
		Select("cb.*").From("character_base as cb").
		Join("accounts as acc").On(" cb.account_id = acc.id").
		Where("cb.server_id = ?", server.ID).
		Where("acc.username = ?", username)


	err = query.One(&character)
	if err != nil {
		logger.BasicLog(err)
		s.conn.Send(packets.BLANK_USER)
		return
	}
	var equips []CharacterEquip

	//ebyte := make([]byte, 60)

	err = s.db.SQL().Select("*").From("character_equip").
		Where("user_id", character.ID).All(&equips)

	if err != nil {
		logger.BasicLog(err)
	}

	fmt.Println(equips)

	//chars := make([]byte, 500)
	//buffer := bytes.NewBuffer(chars)
	//buffer.Write()
	fmt.Println(character.Name)
	s.conn.Send(packets.CHAMPTEST)

	//fmt.Println(character)
	//query, err := s.db.Prepare("SELECT cb.* FROM `character_base` cb " +
	//	"JOIN accounts acc ON cb.account_id = acc.id" +
	//	"JOIN server ON cb.server_id = server.id " +
	//	"WHERE server.identifier = ? " +
	//	"AND acc.username = ? LIMIT 1")
	//if err != nil {
	//	logger.BasicLog(err)
	//	s.conn.Send(packets.BLANK)
	//}
	//
	//row, err := query.Query(username, identifier)
	//if err != nil {
	//	logger.BasicLog(err)
	//	s.conn.Disconnect()
	//}
	//
	//switch err := row.Scan(&account.Username, &account.Password); err {
	//case nil:
	//	s.conn.Send(packets.SERVER_LIST)
	//	break
	//default:
	//	s.LoginFail(3)
	//}
	//
	//defer query.Close()

}

func (s service) CreateAccountCharacter(packet *decode.Decode) {
	fmt.Println(packet.OPCode)
	fmt.Println(*s.conn.ServerID)
}
