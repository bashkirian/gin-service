package main

import (
	"fmt"
	"github.com/bashkirian/gin-service/models"
	"github.com/pressly/goose"
	"os"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func run() error {
	port, err := strconv.Atoi(os.Getenv("db_port"))
	if err != nil {
		return fmt.Errorf("cant parse port: %w", err)
	}

	dbConfig := models.ConnectionConfig{
		Host:     os.Getenv("db_host"),
		Port:     port,
		User:     os.Getenv("db_user"),
		Password: os.Getenv("db_password"),
		DBName:   os.Getenv("db_name"),
	}

	if err := models.ConnectDatabase(dbConfig); err != nil {
		return fmt.Errorf("conn db: %w", err)
	}

	if err := goose.Run("up", models.DB, "migration_db"); err != nil {
		return err
	}

	return nil
}
