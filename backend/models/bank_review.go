package models

type Review struct {
	ID     int   `json:"id"`
	BankID int   `json:"bank_id"`
	Content string `json:"content"`
}

type ReviewPost struct {
	Content string `json:"content"`
}
