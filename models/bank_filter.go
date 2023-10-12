package models

import (
	"github.com/jinzhu/gorm"
)

type BankFilter struct {
	gorm.Model
	Nonstop	bool `json:"nonstop"`
	Self  bool `json:"self"`
	ClientType []string `json:"client_type"`
	Limit int `json:"limit"`
}
