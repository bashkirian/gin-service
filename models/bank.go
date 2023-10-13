package models

type Bank struct {
	ID    int   `json:"id"`
	Name  string `json:"name"`
	Lat string `json:"lat"`
	Lan string `json:"lon"`
}
