package world

type SkillInfo struct {
	ID           int
	Target       int
	Distance     float32
	Range        float32
	MP           int
	CastTime     int
	CoolDown     int
	Pow1         int
	Pow2         int
	Pow3         int
	EquipRequest int
	IsReady      bool
}

func NewSkillInfo() SkillInfo {
	return SkillInfo{}
}
