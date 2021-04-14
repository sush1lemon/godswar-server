package account

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/upper/db/v4"
	"godswar/pkg/decode"
	"godswar/pkg/defaults"
	"godswar/pkg/logger"
	"godswar/pkg/networking"
	"godswar/pkg/packets"
	"godswar/pkg/requests"
	"godswar/pkg/types"
	"godswar/pkg/types/account"
	"godswar/pkg/utility"
	"regexp"
	"strconv"
	"strings"
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
	var request requests.AuthRequest
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
	var request requests.GetCharacterRequest
	var character account.CharacterBase
	var server Server
	reader := bytes.NewReader(packet.Buffer[4:])
	err := binary.Read(reader, binary.LittleEndian, &request)
	if err != nil {
		logger.BasicLog(err)
		s.conn.Disconnect()
	}

	username := utility.RemoveBlank(string(request.Username[:]))
	identifier := string(request.ServerIdentifier[:])

	err = s.db.SQL().Select("*").From("server").
		Where("identifier", identifier).One(&server)

	if err != nil {
		logger.BasicLog("Server not found", err)
		s.conn.Disconnect()
	}
	s.conn.ServerID = &server.ID

	userTbl := s.db.Collection("accounts")
	uQuery := userTbl.Find(db.Cond{"username": username})
	_, err = uQuery.Count()
	if err != nil {
		logger.BasicLog("User not found", err)
		s.conn.Disconnect()
	}

	var user account.Account
	err = uQuery.One(&user)
	if err != nil {
		logger.BasicLog(err)
	}

	s.conn.AccountInfo = &user
	query := s.db.SQL().
		Select("cb.*").From("character_base as cb").
		Where("cb.server_id = ?", server.ID).
		Where("cb.account_id = ?", user.ID)

	err = query.One(&character)
	if err != nil {
		logger.BasicLog(err)
		s.conn.Send(packets.BLANK_USER)
		return
	}

	s.conn.Character = &character
	var equips []account.CharacterEquip
	err = s.db.SQL().Select("*").From("character_equip").
		Where("user_id", character.ID).All(&equips)

	if err != nil {
		logger.BasicLog(err)
	}

	var kitbag account.CharacterKitBag
	err = s.db.SQL().Select("*").From("character_kitbag").
		Where("user_id", character.ID).One(&kitbag)
	if err != nil {
		logger.BasicLog("Cannot fetch character kitbag", err)
	}

	go s.generateUserInfo(character, kitbag)
	equip := *kitbag.Equip

	_, _, equipIDs := s.parseKitBag(equip, true, false)
	preview := s.generateUserPreview(character, equipIDs)
	previewBytes := preview.Bytes()
	s.conn.PreviewBytes = &previewBytes
	s.conn.Send(previewBytes)
}

func (s service) CreateAccountCharacter(packet *decode.Decode) {
	var request requests.CreateCharacterRequest
	reader := bytes.NewReader(packet.DecodedBuffer[:])
	err := binary.Read(reader, binary.LittleEndian, &request)
	if err != nil {
		logger.BasicLog(err)
		s.conn.Disconnect()
	}

	name := utility.RemoveBlank(string(request.CharName[:]))
	var gender string
	var curMap int
	if request.Gender == 0 {
		gender = "female"
	} else {
		gender = "male"
	}

	if request.Camp == 1 {
		curMap = 1
	} else {
		curMap = 2
	}
	

	cb := account.CharacterBase{
		AccountID:           s.conn.AccountInfo.ID,
		ServerID:            *s.conn.ServerID,
		Name:                name,
		Gender:              gender,
		GM:                  0,
		Camp:                int(request.Camp),
		Profession:          int(request.Class),
		FighterJobLv:        1,
		ScholarJobLv:        0,
		FighterJobExp:       0,
		ScholarJobExp:       0,
		CurHP:               1500,
		CurMP:               177,
		Status:              0,
		Belief:              int(request.Faith),
		Prestige:            0,
		Consortia:           0,
		ConsortiaJob:        0,
		ConsortiaContribute: 0,
		StoreNum:            0,
		BagNum:              0,
		HairStyle:           int(request.Hair),
		FaceShape:           int(request.Face),
		CurrentMap:          curMap,
		PosX:                165.00,
		PosZ:                -97.00,
		Money:               10000,
		Stone:               10,
		SkillPoint:          10,
		SkillExp:            0,
		MaxHP:               1500,
		MaxMP:               177,
	}

	r, err :=s.db.SQL().InsertInto("character_base").Values(cb).Exec()
	if err != nil {
		logger.BasicLog("ERROR INSERINT", err)
		s.conn.Send(defaults.USERNAME_TAKEN)
		return
	}

	var defEquip string
	switch request.Class {
	case 0:
		defEquip = defaults.DEFAULT_EQUIPWR
		break
	case 1:
		defEquip = defaults.DEFAULT_EQUIPCH
		break
	case 2:
		defEquip = defaults.DEFAULT_EQUIPPR
		break
	case 3:
		defEquip = defaults.DEFAULT_EQUIPMG
		break
	}

	id, err :=  r.LastInsertId()
	if err != nil {
		logger.BasicLog("Cannot get last insert id")
	}

	kitBag := defaults.DEFAULT_KITBAG

	kitBagData := account.CharacterKitBag{
		CharID: int(id),
		KitBag1: kitBag,
		KitBag2: nil,
		KitBag3: nil,
		KitBag4: nil,
		Storage: nil,
		Equip:   &defEquip,
	}

	_, err = s.db.Collection("character_kitbag").Insert(kitBagData)
	if err != nil {
		logger.BasicLog("Error in creating kitbag", err)
	}

	s.generateUserInfo(cb, kitBagData)
	_, _, equipIDs := s.parseKitBag(defEquip, true, false)
	s.generateUserPreview(cb, equipIDs)
	s.conn.Send([]byte{0x0C, 0x00, 0xB4, 0x27, 0x13, 0x27, 0x8D, 0x0B, 0x01, 0x00, 0x00, 0x00})
}

func (s service) DeleteAccountCharacter(packet *decode.Decode)  {

}

func (s service) parseKitBag(bag string, getEmpty bool, emptyOnly bool) (bytes.Buffer, map[int]types.Item,  bytes.Buffer) {

	var kitbag bytes.Buffer
	var kitbagids bytes.Buffer
	kitbagArr := make(map[int]types.Item)
	kbd := regexp.MustCompile(`#`)
	eqd := regexp.MustCompile(`,`)

	items := kbd.Split(bag, -1)
	for i, e := range items {
		if e == "[]" && getEmpty {
			e = "[0,,,,,,1,1,0,1,0]"
		}

		if e == "[]" {
			if getEmpty {
				padding := make([]byte, 72)
				kitbag.Write(padding)
				kitbagids.Write([]byte{0x00, 0x00, 0x00, 0x00})
			}
		} else if e != "" {
			e = strings.Replace(e, "[", "", -1)
			e = strings.Replace(e, "]", "", -1)
			item := eqd.Split(e, -1)
			itembar := make([]byte, 72)
			itemBuff := bytes.NewBuffer(itembar)
			itemBuff.Reset()

			id := s.equipToData(item[0])

			if emptyOnly && item[0] != "0" {
				continue
			}

			s1 := s.equipToData(item[1])
			s2 := s.equipToData(item[2])
			s3 := s.equipToData(item[3])
			s4 := s.equipToData(item[4])
			s5 := s.equipToData(item[5])

			quality := uint8(s.stringToInt(item[6]))
			grade := uint8(s.stringToInt(item[7]))
			bound := uint8(s.stringToInt(item[8]))
			stack := uint8(s.stringToInt(item[9]))
			exp := s.equipToData(item[10])
			padding := make([]byte, 36)

			//rand.Seed(time.Now().UnixNano())
			//token := make([]byte, 4)
			//rand.Read(token)
			//token[3] = byte(i)

			itemBuff.Write(id)
			itemBuff.Write(s1)
			itemBuff.Write(s2)
			itemBuff.Write(s3)
			itemBuff.Write(s4)
			itemBuff.Write(s5)
			itemBuff.WriteByte(quality)
			itemBuff.WriteByte(grade)
			itemBuff.WriteByte(bound)
			itemBuff.WriteByte(stack)
			itemBuff.Write(exp)
			itemBuff.Write(padding)

			if i == len(items) {
				itemBuff.Write([]byte{0x00, 0x00, 0x00, 0x00})
			} else {
				itemBuff.Write([]byte{0x42, 0x00, 0x00, 0x00})
			}

			itemBuff.WriteTo(&kitbag)
			kitbagids.Write(id)

			if getEmpty && !emptyOnly {
				kitbagArr[i] = types.Item{
					ID:      id,
					Attr1:   s1,
					Attr2:   s2,
					Attr3:   s3,
					Attr4:   s4,
					Attr5:   s5,
					Quality: quality,
					Grade:   grade,
					Bound:   bound,
					Stack:   stack,
					Exp:     exp,
					Unk:     padding,
					Ending: []byte{0x42, 0x00, 0x00, 0x00},
				}
			}
		}
	}

	return kitbag, kitbagArr, kitbagids
}

func (s service) equipToData(e string) []byte {
	if e == "" {
		return []byte{0xFF, 0xFF, 0xFF, 0xFF}
	}

	id := s.stringToInt(e)
	eq := make([]byte, 4)
	binary.LittleEndian.PutUint16(eq, uint16(id))
	return eq
}

func (s service) stringToInt(e string) int {
	id, err := strconv.Atoi(e)
	if err != nil {
		logger.BasicLog("CANNOT PARSE STRING", e, err)
	}
	return id
}
