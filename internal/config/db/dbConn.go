package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
	"go.uber.org/zap"
)

func InitPool(ctx context.Context) *pgxpool.Pool {
	logger.Debug("Initializing database connection pool")

	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	dbname := os.Getenv("DATABASE_NAME")
	user := os.Getenv("DATABASE_USER")

	logger.Debug("Database connection parameters",
		zap.String("host", host),
		zap.String("port", port),
		zap.String("database", dbname),
		zap.String("user", user))

	pool, err := pgxpool.New(ctx, fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		os.Getenv("DATABASE_PASSWORD"),
		host,
		port,
		dbname))
	if err != nil {
		logger.Error("Failed to create database connection pool", err)
		panic(err)
	}

	// Test the connection
	if err := pool.Ping(ctx); err != nil {
		logger.Error("Failed to ping database", err)
		panic(err)
	}

	logger.Info("Database connection pool initialized successfully")
	return pool
}
