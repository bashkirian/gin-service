package models

import (
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	ID	uint `json:"id" binding:"required" gorm:"primary_key"`
	LocationLat string `json:"location_lat"`
	LocationLon string `json:"location_lon"`
}
