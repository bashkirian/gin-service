package models

type BankFilter struct {
	Nonstop	bool `json:"nonstop"`
	Self  bool `json:"self"`
	ClientType []string `json:"client_type"`
	Limit int `json:"limit"`
}
