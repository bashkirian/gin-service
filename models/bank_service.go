package models

import (
	"github.com/jinzhu/gorm"
)

type BankService struct {
	gorm.Model
	BankServiceID int `gorm:"primary_key"`
	BankID     int
	ServiceID  int
	Service Service
	Bank Bank
}
