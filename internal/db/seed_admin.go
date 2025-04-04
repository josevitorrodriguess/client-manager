package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/josevitorrodriguess/client-manager/internal/db/sqlc"
	"golang.org/x/crypto/bcrypt"
)

func CreateAdmin(ctx context.Context, pool *pgxpool.Pool) error {
	adminName := os.Getenv("ADMIN_NAME")
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin := sqlc.CreateUserParams{
		Name:     adminName,
		Email:    adminEmail,
		Password: string(hashedPassword),
		IsAdmin:  true,
	}

	_, err = sqlc.New(pool).CreateUser(ctx, admin)
	if err != nil {
		return err
	}

	return nil
}
