package models

type Bank struct {
	ID              int
	Name            string  `json:"salePointName" `
	Address         string  `json:"address"`
	Status          string  `json:"status"`
	Rko             string  `json:"rko"`
	OfficeType      string  `json:"officeType"`
	SalePointFormat string  `json:"salePointFormat"`
	SuoAvailability string  `json:"suoAvailability"`
	HasRamp         string  `json:"hasRamp"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	MetroStation    string  `json:"metroStation"`
	Distance        int64   `json:"distance"`
	Kep             bool    `json:"kep"`
	MyBranch        bool    `json:"myBranch"`
}
