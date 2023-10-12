package models

type Service struct {
	id uint `gorm:"primary_key"`
	description string
}
