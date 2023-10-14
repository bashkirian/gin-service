package main

import (
	"fmt"
	"github.com/bashkirian/gin-service/repo"
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

	dbConfig := repo.ConnectionConfig{
		Host:     os.Getenv("db_host"),
		Port:     port,
		User:     os.Getenv("db_user"),
		Password: os.Getenv("db_password"),
		DBName:   os.Getenv("db_name"),
	}

	if err := repo.ConnectDatabase(dbConfig); err != nil {
		return fmt.Errorf("conn db: %w", err)
	}

	if err := goose.Run("up", repo.Storage(), "migration_db"); err != nil {
		return err
	}

	return nil
}
