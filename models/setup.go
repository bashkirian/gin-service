package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("postgres", "port=5432 user=postgres password=12345 dbname=goservice sslmode=disable")

	if err != nil {
		fmt.Println(err)
	}
	database.AutoMigrate(&Bank{})
	bank := Bank{Name: "Vernadsky", 
	LocationLat: "30.0", 
	LocationLan: "20.0" }
	database.Create(&bank)
	database.AutoMigrate(&Review{}, &Client{}, &Service{}, &BankService{}, &Review{}, &Client{})
	database.Model(&BankService{}).AddForeignKey("bank_id", "banks(id)", "RESTRICT", "RESTRICT")
	
	DB = database
}
