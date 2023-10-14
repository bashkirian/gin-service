package models

import (
	"github.com/google/uuid"
)
type Bank struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"salePointName" db:"salePointName"`
	Address string `json:"address" db:"address"`
	Status string `json:"status" db:"status"`
	Rko string `json:"rko" db:"rko"`
	OfficeType string `json:"officeType" db:"officeType"`
	SalePointFormat string `json:"salePointFormat" db:"salePointFormat"`
	SuoAvailability string `json:"suoAvailability" db:"suoAvailability"`
	HasRamp string `json:"hasRamp" db:"hasRamp"`
	Latitude float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
	MetroStation string `json:"metroStation" db:"metroStation"`
	Distance int64 `json:"distance" db:"distance"`
	Kep bool `json:"kep" db:"kep"`
	MyBranch bool `json:"myBranch" db:"myBranch"`
}
