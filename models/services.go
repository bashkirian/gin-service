package models

import (
	"github.com/jinzhu/gorm"
)

type Service struct {
	gorm.Model
	ID int `gorm:"primary_key"`
	Description string
}
