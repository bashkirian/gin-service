package models

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "log"
	"os"
)

func MigrateDatabase() error {
	var err error
	offices, err := os.ReadFile("migrations/offices.txt")
	if err != nil {
		return fmt.Errorf("readfile: %w", err)
	}
	db, err := sqlx.Connect("postgres", "host = localhost user=postgres dbname=goservice password=12345 sslmode=disable")
	if err != nil {
		return fmt.Errorf("postgres connection: %w", err)
	}
	var banks []Bank
	if err = json.Unmarshal([]byte(offices), &banks); err != nil {
		return fmt.Errorf("json unmarshal: %w", err)
	}
	query := `INSERT INTO bank.banks(salePointName, address, status, rko, officeType, salePointFormat, suoAvailability, hasRamp, latitude, longitude, metroStation, distance, kep, myBranch) 
          VALUES(:salePointName, :address, :status, :rko, :officeType, :salePointFormat, :suoAvailability, :hasRamp, :latitude, :longitude, :metroStation, :distance, :kep, :myBranch)
		  ON CONFLICT DO NOTHING`

	for _, bank := range banks {
		_, err := db.NamedExec(query, bank)
		if err != nil {
			return fmt.Errorf("named exec: %w", err)
		}
	}

	return nil
}
