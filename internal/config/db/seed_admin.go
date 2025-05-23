package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
	"github.com/josevitorrodriguess/client-manager/internal/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(ctx context.Context, pool *pgxpool.Pool) error {
	adminName := os.Getenv("ADMIN_NAME")
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	queries := sqlc.New(pool)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("Failed to hash admin password", err)
		return err
	}

	admin := sqlc.CreateUserParams{
		Name:     adminName,
		Email:    adminEmail,
		Password: string(hashedPassword),
		IsAdmin:  true,
	}

	_, err = queries.CreateUser(ctx, admin)
	if err != nil {
		logger.Error("Failed to create admin user", err)
		return err
	}

	return nil
}
