package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_NAME     = "DB_NAME"
)

func InitPool(ctx context.Context) *pgxpool.Pool {
	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=% dbname=%s  sslmode=disable",
		"DB_USER",
		"DB_PASSWORD",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME"))
	if err != nil {
		panic(err)
	}

	return pool
}
