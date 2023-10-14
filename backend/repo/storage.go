package repo

import (
	"database/sql"
	"fmt"
	"time"
)

var db *sql.DB

type ConnectionConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func ConnectDatabase(conf ConnectionConfig) error {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)

	time.Sleep(1 * time.Second)

	// open database
	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}

	// check DB
	if err = db.Ping(); err != nil {
		return fmt.Errorf("ping: %w", err)
	}

	return nil
}
