package models

type Bank struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	LocationLat uint 
	LocationLan uint
}
