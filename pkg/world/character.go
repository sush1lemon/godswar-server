package world

type CharacterBase struct {
	Name          [32]byte
	Gender        [1]byte
	GM            [1]byte
	Camp          [1]byte
	Profession    [1]byte
	FighterJobLv  [1]byte
	ScholarJobLv  [1]byte
	FighterJobExp [4]byte
	ScholarJobExp [4]byte
	CurrentHP     [16]byte

}
