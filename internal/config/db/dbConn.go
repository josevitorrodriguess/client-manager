package db

import (
	"database/sql"

	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
)

const (
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_NAME     = "DB_NAME"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("postgres", "user="+DB_USER+" password="+DB_PASSWORD+" host="+DB_HOST+" port="+DB_PORT+" dbname="+DB_NAME+" sslmode=disable")
	if err != nil {
		logger.Error("Error connecting to database: ", err)
		panic(err)
	}

	defer db.Close()

	logger.Info("Connected to database")
	return db
}
