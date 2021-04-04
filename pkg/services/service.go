package services

import (
	"database/sql"
	"godswar/pkg/networking"
	"godswar/pkg/services/account"
)


type Service struct {
	Account account.Service
}

func NewService(db *sql.DB, conn *networking.Connection) *Service {
	return &Service{
		Account: account.NewAccountService(db, conn),
	}
}