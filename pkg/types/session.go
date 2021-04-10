package types

import "godswar/pkg/types/account"

type UserSession struct {
	AccountInfo    *account.Account
	Character      *account.CharacterBase
	Equip          *map[int]Item
	PreviewBytes   *[]byte
	EquipBytes     *[]byte
	CharacterBytes *[]byte
	Bag1           *map[int]Item
	Bag2           *map[int]Item
}
