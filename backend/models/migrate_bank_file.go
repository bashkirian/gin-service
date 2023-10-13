package models

import (
    "encoding/json"
    _ "fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func MigrateDatabase() {
	offices, err := os.ReadFile("migrations/offices.txt")
	if err != nil {
		log.Fatalln(err)
	}
	db, err := sqlx.Connect("postgres", "host = localhost user=postgres dbname=goservice password=12345 sslmode=disable")
    if err != nil {
		log.Fatalln(err)
	}
	var banks []Bank
    if err = json.Unmarshal([]byte(offices), &banks); err != nil {
        panic(err)
    }
    query := `INSERT INTO banks(salePointName, address, status, rko, officeType, salePointFormat, suoAvailability, hasRamp, latitude, longitude, metroStation, distance, kep, myBranch) 
          VALUES(:salePointName, :address, :status, :rko, :officeType, :salePointFormat, :suoAvailability, :hasRamp, :latitude, :longitude, :metroStation, :distance, :kep, :myBranch)
		  ON CONFLICT DO NOTHING`

	for _, bank := range banks {
		_, err := db.NamedExec(query, bank)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
