package models

import (
    "encoding/json"
    _ "fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

var str = `[
    {
        "salePointName": "ДО «Солнечногорский» Филиала № 7701 Банка ВТБ (ПАО)",
        "address": "141506, Московская область, г. Солнечногорск, ул. Красная, д. 60",        "status": "открытая",
        "openHours": [
            {
                "days": "пн",
                "hours": "09:00-18:00"
            },
            {
                "days": "вт",
                "hours": "09:00-18:00"
            },
            {
                "days": "ср",
                "hours": "09:00-18:00"
            },
            {
                "days": "чт",
                "hours": "09:00-18:00"
            },
            {
                "days": "пт",
                "hours": "09:00-17:00"
            },
            {
                "days": "сб",
                "hours": "выходной"
            },
            {
                "days": "вс",
                "hours": "выходной"
            }
        ],
        "rko": "есть РКО",        "openHoursIndividual": [
            {
                "days": "пн",
                "hours": "09:00-20:00"
            },
            {
                "days": "вт",
                "hours": "09:00-20:00"
            },
            {
                "days": "ср",
                "hours": "09:00-20:00"
            },
            {
                "days": "чт",
                "hours": "09:00-20:00"
            },
            {
                "days": "пт",
                "hours": "09:00-20:00"
            },
            {
                "days": "сб",
                "hours": "10:00-17:00"
            },
            {
                "days": "вс",
                "hours": "выходной"
            }
        ],
        "officeType": "Да (Зона Привилегия)",
        "salePointFormat": "Универсальный",
        "suoAvailability": "Y",
        "hasRamp": "N",
        "latitude": 56.184479,
        "longitude": 36.984314,
        "metroStation": null,
        "distance": 62105,
        "kep": true,
        "myBranch": false
    }
]`

type BankSql struct {
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

func MigrateDatabase() {
	db, err := sqlx.Connect("postgres", "host = localhost user=postgres dbname=goservice password=12345 sslmode=disable")
    if err != nil {
		log.Fatalln(err)
	}
	var banks []BankSql
    if err = json.Unmarshal([]byte(str), &banks); err != nil {
        panic(err)
    }
    query := `INSERT INTO banks(salePointName, address, status, rko, officeType, salePointFormat, suoAvailability, hasRamp, latitude, longitude, metroStation, distance, kep, myBranch) 
          VALUES(:salePointName, :address, :status, :rko, :officeType, :salePointFormat, :suoAvailability, :hasRamp, :latitude, :longitude, :metroStation, :distance, :kep, :myBranch)`

	for _, bank := range banks {
		_, err := db.NamedExec(query, bank)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
