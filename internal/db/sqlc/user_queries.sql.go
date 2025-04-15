// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user_queries.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const checkIfUserIsAdmin = `-- name: CheckIfUserIsAdmin :one
SELECT 
  is_admin
FROM users
WHERE id = $1
`

func (q *Queries) CheckIfUserIsAdmin(ctx context.Context, id uuid.UUID) (bool, error) {
	row := q.db.QueryRow(ctx, checkIfUserIsAdmin, id)
	var is_admin bool
	err := row.Scan(&is_admin)
	return is_admin, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (name, email, password, is_admin) 
VALUES ($1, $2, $3, $4)
RETURNING id
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.IsAdmin,
	)
	var id uuid.UUID
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT 
    id,
    name,
    email,
    password,
    created_at,
    updated_at
FROM users
WHERE email = $1
`

type GetUserByEmailRow struct {
	ID        uuid.UUID          `json:"id"`
	Name      string             `json:"name"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name = $1, email = $2, password = $3, is_admin = $4
WHERE id = $4
RETURNING id, name, email
`

type UpdateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type UpdateUserRow struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.Name,
		arg.Email,
		arg.Password,
		arg.IsAdmin,
	)
	var i UpdateUserRow
	err := row.Scan(&i.ID, &i.Name, &i.Email)
	return i, err
}
