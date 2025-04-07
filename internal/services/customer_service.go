package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
	"github.com/josevitorrodriguess/client-manager/internal/db/sqlc"
	"github.com/josevitorrodriguess/client-manager/internal/validators/customer"
	"go.uber.org/zap"
)

var ErrDuplicatedData = errors.New("cpf, phone or email already exists")

type CustomerService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewCustomerService(pool *pgxpool.Pool) *CustomerService {
	return &CustomerService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (cs *CustomerService) CreatePFCustomer(ctx context.Context, customer customer.CustomerPFRequest) (uuid.UUID, error) {
	args := sqlc.CreateCustomerPFParams{
		Type:        sqlc.CustomerType(customer.Type),
		Email:       customer.Email,
		Phone:       customer.Phone,
		Cpf:         customer.Cpf,
		Name:        customer.Name,
		BirthDate:   customer.BirthDate,
		AddressType: customer.AddressType,
		Street:      customer.Street,
		Complement:  customer.Complement,
		State:       customer.State,
		City:        customer.City,
		Cep:         customer.Cep,
	}

	id, err := cs.queries.CreateCustomerPF(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicatedData
		}
		logger.Error("Failed to create PF customer", err,
			zap.String("name", customer.Name),
			zap.String("email", customer.Email))
		return uuid.UUID{}, err
	}

	return id, nil
}

func (cs *CustomerService) CreatePJCustomer(ctx context.Context, customer customer.CustomerPJRequest) (uuid.UUID, error) {
	args := sqlc.CreateCustomerPJParams{
		Type:        sqlc.CustomerType(customer.Type),
		Email:       customer.Email,
		Phone:       customer.Phone,
		Cnpj:        customer.Cnpj,
		CompanyName: customer.CompanyName,
		AddressType: customer.AddressType,
		Street:      customer.Street,
		Complement:  customer.Complement,
		State:       customer.State,
		City:        customer.City,
		Cep:         customer.Cep,
	}

	id, err := cs.queries.CreateCustomerPJ(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return uuid.UUID{}, ErrDuplicatedData
		}
		logger.Error("Failed to create PJ customer", err,
			zap.String("company_name", customer.CompanyName),
			zap.String("email", customer.Email))
		return uuid.UUID{}, err
	}

	return id, nil
}

func (cs *CustomerService) AddAddressRequest(ctx context.Context, address customer.AddAddressRequest) (int32, error) {
	args := sqlc.AddAddressToCustomerParams{
		CustomerID:  address.CustomerID,
		AddressType: address.AddressType,
		Street:      address.Street,
		Number:      address.Number,
		Complement:  address.Complement,
		State:       address.State,
		City:        address.City,
		Cep:         address.Cep,
	}

	id, err := cs.queries.AddAddressToCustomer(ctx, args)
	if err != nil {
		logger.Error("Failed to add address to customer", err,
			zap.String("customer_id", address.CustomerID.String()))
		return 0, err
	}

	return id, nil
}
