package models

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "log"
	"os"
)

func PopulateDatabase() error {
	var err error
	offices, err := os.ReadFile("migration_db/offices.txt")
	if err != nil {
		return err
	}
	db, err := sqlx.Connect("postgres", "host = db user=root dbname=defaultdb password=1 sslmode=disable")
	if err != nil {
		return err
	}
	var banks_temp []BankTemp
	if err = json.Unmarshal([]byte(offices), &banks_temp); err != nil {
		panic(err)
	}

	var banks []Bank
	for _, bank_temp := range banks_temp {
		var bank Bank
		bank.ID = bank_temp.ID
		bank.Name = bank_temp.Name
		bank.Address = bank_temp.Address
		bank.Status = bank_temp.Status
		if (bank_temp.Rko == "Есть РКО") {
			bank.Rko = true
		} else if (bank_temp.Rko == "Нет РКО") {
			bank.Rko = false
		}
		bank.OfficeType = bank_temp.OfficeType
		bank.SalePointFormat = bank_temp.SalePointFormat
		if (bank_temp.SuoAvailability == "Y") {
			bank.SuoAvailability = true
		} else if (bank_temp.SuoAvailability == "N") {
			bank.SuoAvailability = false
		}
		if (bank_temp.HasRamp == "Y") {
			bank.HasRamp = true
		} else if (bank_temp.HasRamp == "N") {
			bank.HasRamp = false
		}
		bank.Latitude = bank_temp.Latitude
		bank.Longitude = bank_temp.Longitude
		bank.MetroStation = bank_temp.MetroStation
		bank.Distance = bank_temp.Distance
		bank.Kep = bank_temp.Kep
		bank.MyBranch = bank_temp.MyBranch
		banks = append(banks, bank)
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
