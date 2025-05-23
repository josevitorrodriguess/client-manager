// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: services_queries.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const countServicesByCustomerID = `-- name: CountServicesByCustomerID :one
SELECT COUNT(*) FROM services
WHERE customer_id = $1
`

func (q *Queries) CountServicesByCustomerID(ctx context.Context, customerID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, countServicesByCustomerID, customerID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createService = `-- name: CreateService :one
INSERT INTO services (
    customer_id,
    type_product,
    description,
    total_value,
    down_payment,
    is_paid,
    is_finished
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) 
RETURNING id
`

type CreateServiceParams struct {
	CustomerID  uuid.UUID      `json:"customer_id"`
	TypeProduct string         `json:"type_product"`
	Description string         `json:"description"`
	TotalValue  pgtype.Numeric `json:"total_value"`
	DownPayment pgtype.Numeric `json:"down_payment"`
	IsPaid      bool           `json:"is_paid"`
	IsFinished  bool           `json:"is_finished"`
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) (int32, error) {
	row := q.db.QueryRow(ctx, createService,
		arg.CustomerID,
		arg.TypeProduct,
		arg.Description,
		arg.TotalValue,
		arg.DownPayment,
		arg.IsPaid,
		arg.IsFinished,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const deleteService = `-- name: DeleteService :exec
DELETE FROM services
WHERE id = $1
`

func (q *Queries) DeleteService(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteService, id)
	return err
}

const getServicesByCustomerID = `-- name: GetServicesByCustomerID :many
SELECT id, customer_id, type_product, description, total_value, down_payment, is_paid, is_finished FROM services
WHERE customer_id = $1
ORDER BY id
`

func (q *Queries) GetServicesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]Service, error) {
	rows, err := q.db.Query(ctx, getServicesByCustomerID, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Service
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.TypeProduct,
			&i.Description,
			&i.TotalValue,
			&i.DownPayment,
			&i.IsPaid,
			&i.IsFinished,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAllServices = `-- name: ListAllServices :many
SELECT id, customer_id, type_product, description, total_value, down_payment, is_paid, is_finished FROM services
ORDER BY id
`

func (q *Queries) ListAllServices(ctx context.Context) ([]Service, error) {
	rows, err := q.db.Query(ctx, listAllServices)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Service
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ID,
			&i.CustomerID,
			&i.TypeProduct,
			&i.Description,
			&i.TotalValue,
			&i.DownPayment,
			&i.IsPaid,
			&i.IsFinished,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateServiceFinishStatus = `-- name: UpdateServiceFinishStatus :one
UPDATE services
SET is_finished = $1
WHERE id = $2
RETURNING id, customer_id, type_product, description, total_value, down_payment, is_paid, is_finished
`

type UpdateServiceFinishStatusParams struct {
	IsFinished bool  `json:"is_finished"`
	ID         int32 `json:"id"`
}

func (q *Queries) UpdateServiceFinishStatus(ctx context.Context, arg UpdateServiceFinishStatusParams) (Service, error) {
	row := q.db.QueryRow(ctx, updateServiceFinishStatus, arg.IsFinished, arg.ID)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.TypeProduct,
		&i.Description,
		&i.TotalValue,
		&i.DownPayment,
		&i.IsPaid,
		&i.IsFinished,
	)
	return i, err
}

const updateServicePaymentStatus = `-- name: UpdateServicePaymentStatus :one
UPDATE services
SET is_paid = $1
WHERE id = $2
RETURNING id, customer_id, type_product, description, total_value, down_payment, is_paid, is_finished
`

type UpdateServicePaymentStatusParams struct {
	IsPaid bool  `json:"is_paid"`
	ID     int32 `json:"id"`
}

func (q *Queries) UpdateServicePaymentStatus(ctx context.Context, arg UpdateServicePaymentStatusParams) (Service, error) {
	row := q.db.QueryRow(ctx, updateServicePaymentStatus, arg.IsPaid, arg.ID)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.CustomerID,
		&i.TypeProduct,
		&i.Description,
		&i.TotalValue,
		&i.DownPayment,
		&i.IsPaid,
		&i.IsFinished,
	)
	return i, err
}
