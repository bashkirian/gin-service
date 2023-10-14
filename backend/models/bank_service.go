package models

import "github.com/google/uuid"

type BankService struct {
	BankID    uuid.UUID
	ServiceID uuid.UUID
}
