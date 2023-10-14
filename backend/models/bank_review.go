package models

import "github.com/google/uuid"

type Review struct {
	ID      uuid.UUID `json:"id"`
	BankID  uuid.UUID `json:"bank_id"`
	Content string    `json:"content"`
}

type ReviewPost struct {
	Content string `json:"content"`
}
