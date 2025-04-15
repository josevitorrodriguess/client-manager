// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddAddressToCustomer(ctx context.Context, arg AddAddressToCustomerParams) (int32, error)
	CheckIfUserIsAdmin(ctx context.Context, id uuid.UUID) (bool, error)
	CountServicesByCustomerID(ctx context.Context, customerID uuid.UUID) (int64, error)
	CreateCustomerPF(ctx context.Context, arg CreateCustomerPFParams) (uuid.UUID, error)
	CreateCustomerPJ(ctx context.Context, arg CreateCustomerPJParams) (uuid.UUID, error)
	CreateService(ctx context.Context, arg CreateServiceParams) (int32, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (uuid.UUID, error)
	DeleteAddress(ctx context.Context, id int32) error
	DeleteCustomer(ctx context.Context, id uuid.UUID) error
	DeleteService(ctx context.Context, id int32) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetAllCustomers(ctx context.Context) ([]GetAllCustomersRow, error)
	GetCustomerAddresses(ctx context.Context, customerID uuid.UUID) ([]GetCustomerAddressesRow, error)
	GetCustomerByID(ctx context.Context, id uuid.UUID) (GetCustomerByIDRow, error)
	GetServicesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]Service, error)
	GetUserByEmail(ctx context.Context, email string) (GetUserByEmailRow, error)
	ListAllServices(ctx context.Context) ([]Service, error)
	UpdateAddress(ctx context.Context, arg UpdateAddressParams) (int32, error)
	UpdateCustomerBasicInfo(ctx context.Context, arg UpdateCustomerBasicInfoParams) (uuid.UUID, error)
	UpdateServiceFinishStatus(ctx context.Context, arg UpdateServiceFinishStatusParams) (Service, error)
	UpdateServicePaymentStatus(ctx context.Context, arg UpdateServicePaymentStatusParams) (Service, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
}

var _ Querier = (*Queries)(nil)
