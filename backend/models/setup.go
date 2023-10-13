package models

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "12345"
	dbname = "goservice"
)

var DB *sql.DB

func ConnectDatabase() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
    // open database
	var err error
    DB, err = sql.Open("postgres", psqlconn)
    CheckError(err)
 
    // check DB
    err = DB.Ping()
    CheckError(err)
	
    fmt.Println("Connected!")

}
 
func CheckError(err error) {
    if err != nil {
        log.Fatalln(err)
    }
}
