package models

import (
	"github.com/jinzhu/gorm"
)

type Bank struct {
	gorm.Model
	ID    int   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	LocationLat string `json:"lat"`
	LocationLan string `json:"lon"`
}
