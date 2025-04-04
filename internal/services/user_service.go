package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/josevitorrodriguess/client-manager/internal/db/sqlc"
	"github.com/josevitorrodriguess/client-manager/internal/utils"
	"github.com/josevitorrodriguess/client-manager/internal/validators/user"
)

var (
	ErrDuplicatedEmailOrUsername = errors.New("username or email already exists")
	ErrInvalidCredentials        = errors.New("invalid credentials")
)

type UserService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (us *UserService) Create(ctx context.Context, user user.UserRequest) (uuid.UUID, error) {
	hashPass, err := utils.EncryptPassword(user.Password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error encrypting password: %w", err)
	}

	args := sqlc.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashPass,
		IsAdmin: user.IsAdmin,
	}

	id, err := us.queries.CreateUser(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicatedEmailOrUsername
		}
		return uuid.UUID{}, err
	}
	return id, nil
}

func (us *UserService) AuthenticateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	user, err := us.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, ErrInvalidCredentials
		}
		return uuid.UUID{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return uuid.UUID{}, ErrInvalidCredentials
		}
		return uuid.UUID{}, err
	}
	return user.ID, nil
}

func (us *UserService) CheckIsAdmin(ctx context.Context, id uuid.UUID) (bool, error) {
	ok, err := us.queries.CheckIfUserIsAdmin(ctx, id)
	if err != nil {
		return false, err
	}
	return ok, nil
}
