package account

type Account struct {
	ID          int    `db:"id"`
	Username    string `db:"username"`
	Password    string `db:"password"`
	LoginStatus int    `db:"login_status"`
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
	EarlRank            int     `db:"earl_rank"`
	Consortia           int     `db:"consortia"`
	ConsortiaJob        int     `db:"consortia_job"`
	ConsortiaContribute int     `db:"consortia_contribute"`
	StoreNum            int     `db:"store_num"`
	BagNum              int     `db:"bag_num"`
	HairStyle           int     `db:"hair_style"`
	FaceShape           int     `db:"face_shap"`
	CurrentMap          int     `db:"Map"`
	PosX                float32 `db:"Pos_X"`
	PosZ                float32 `db:"Pos_Z"`
	Money               int     `db:"Money"`
	Stone               int     `db:"Stone"`
	SkillPoint          int     `db:"SkillPoint"`
	SkillExp            int     `db:"SkillExp"`
	MaxHP               int     `db:"MaxHP"`
	MaxMP               int     `db:"MaxMP"`
	RegisterTime        string  `db:"Register_time"`
	LastLoginTime       string  `db:"LastLogin_time"`
	MuteTime            int     `db:"mutetime"`
}

type CharacterEquip struct {
	CharID     int      `db:"user_id"`
	BodyPartID int      `db:"body_part_id"`
	PropID     int      `db:"prop_id"`
	Type1      *int     `db:"type1"`
	Quality1   *int     `db:"quality1"`
	Value1     *float32 `db:"value1"`
	Type2      *int     `db:"type2"`
	Quality2   *int     `db:"quality2"`
	Value2     *float32 `db:"value2"`
	Type3      *int     `db:"type3"`
	Quality3   *int     `db:"quality3"`
	Value3     *float32 `db:"value3"`
	Type4      *int     `db:"type4"`
	Quality4   *int     `db:"quality4"`
	Value4     *float32 `db:"value4"`
	Type5      *int     `db:"type5"`
	Quality5   *int     `db:"quality5"`
	Value5     *float32 `db:"value5"`
	IsBind     int      `db:"isbind"`
}

type CharacterKitBag struct {
	CharID  int     `db:"user_id"`
	KitBag1 string  `db:"kitbag_1"`
	KitBag2 *string `db:"kitbag_2"`
	KitBag3 *string `db:"kitbag_3"`
	KitBag4 *string `db:"kitbag_4"`
	Storage *string `db:"storage"`
	Equip   *string `db:"equip"`
}
