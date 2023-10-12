package models

import (
	"github.com/jinzhu/gorm"
)

type Review struct {
	gorm.Model
	ID     int   `json:"id" gorm:"primary_key"`
	BankID int   `json:"bank_id"`
	Content string `json:"content"`
	Bank Bank
}

type ReviewPost struct {
	Content string `json:"content"`
}
