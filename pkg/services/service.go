package services

import (
	"github.com/upper/db/v4"
	"godswar/pkg/networking"
	"godswar/pkg/services/account"
)

type Service struct {
	Account account.Service
}

func NewService(db db.Session, conn *networking.Connection) *Service {
	return &Service{
		Account: account.NewAccountService(db, conn),
	}
}
