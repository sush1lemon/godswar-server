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
}

type service struct {
	db   db.Session
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
	Username         [32]byte
	ServerIdentifier [25]byte
}

type CharacterBase struct {
	ID                  int     `db:"id"`
	AccountID           int     `db:"account_id"`
	ServerID            int     `db:"server_id"`
	Name                string  `db:"name"`
	Gender              string  `db:"gender"`
	GM                  int     `db:"GM"`
	Camp                int     `db:"camp"`
	Profession          int     `db:"profession"`
	FighterJobLv        int     `db:"fighter_job_lv"`
	ScholarJobLv        int     `db:"scholar_job_lv"`
	FighterJobExp       int     `db:"fighter_job_exp"`
	ScholarJobExp       int     `db:"scholar_job_exp"`
	CurHP               int     `db:"curHP"`
	CurMP               int     `db:"curMP"`
	Status              int     `db:"status"`
	Belief              int     `db:"belief"`
	Prestige            int     `db:"prestige"`
	Consortia           int     `db:"consortia"`
	ConsortiaJob        int     `db:"consortia_job"`
	ConsortiaContribute int     `db:"consortia_contribute"`
	StoreNum            int     `db:"store_num"`
	BagNum              int     `db:"bag_num"`
	HairStyle           int     `db:"hair_style"`
	FaceShape           int     `db:"face_shap"`
	CurrentMap          int     `db:"map"`
	PosX                float32 `db:"Pos_X"`
	PosZ                float32 `db:"Pos_Z"`
	Money               int     `db:"money"`
	Stone               int     `db:"stone"`
	SkillPoint          int     `db:"SkillPoint"`
	SkillExp            int     `db:"SkillExp"`
	MaxHP               int     `db:"MaxHP"`
	MaxMP               int     `db:"MaxMP"`
	RegisterTime        string  `db:"Register_time"`
	LastLoginTime       string  `db:"LastLogin_time"`
	MuteTime            string  `db:"mutetime"`
}

type CharacterEquip struct {
	CharID     int
	BodyPartID int
	PropID     int
	Type1      int
	Quality1   int
	Value1     float32
	Type2      int
	Quality2   int
	Value2     float32
	Type3      int
	Quality3   int
	Value3     float32
	Type4      int
	Quality4   int
	Value4     float32
	Type5      int
	Quality5   int
	Value5     float32
	IsBind     int
}


type Server struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	IP          string `db:"ip_address"`
	Identifier  string `db:"identifier"`
	ServerLimit int    `db:"server_limit"`
}