package models

type Review struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	BankID uint   `json:"bank_id"`
	Content string `json:"content"`
}

type ReviewPost struct {
	BankID uint `json:"bank_id"`
	Content string `json:"content"`
}
