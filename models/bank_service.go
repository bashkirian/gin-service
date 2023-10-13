package models

type BankService struct {
	BankID     int
	ServiceID  int
	Service Service
	Bank Bank
}
