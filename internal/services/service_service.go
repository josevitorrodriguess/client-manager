package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
	"github.com/josevitorrodriguess/client-manager/internal/db/sqlc"
	"github.com/josevitorrodriguess/client-manager/internal/validators/service"
)

type ServiceService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewServiceService(pool *pgxpool.Pool) *ServiceService {
	return &ServiceService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

func (ss *ServiceService) CreateService(ctx context.Context, service service.ServiceRequest) (int32, error) {
	data := sqlc.CreateServiceParams{
		CustomerID:  service.CustomerID,
		TypeProduct: service.TypeProduct,
		Description: service.Description,
		TotalValue:  service.TotalValue,
		DownPayment: service.DownPayment,
		IsPaid:      service.IsPaid,
		IsFinished:  service.IsFinished,
	}

	serviceID, err := ss.queries.CreateService(ctx, data)
	if err != nil {
		logger.Error("Failed to create service to customer", err)
		return 0, err
	}
	return serviceID, nil
}

func (ss *ServiceService) GetServicesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]sqlc.Service, error) {
	services, err := ss.queries.GetServicesByCustomerID(ctx, customerID)
	if err != nil {
		logger.Error("Failed to get services for customer", err)
		return nil, err
	}
	return services, nil
}

func (ss *ServiceService) CountServicesByCustomerID(ctx context.Context, customerID uuid.UUID) (int64, error) {
	count, err := ss.queries.CountServicesByCustomerID(ctx, customerID)
	if err != nil {
		logger.Error("Failed to count services for customer", err)
		return 0, err
	}
	return count, nil
}

func (ss *ServiceService) ListAllServices(ctx context.Context) ([]sqlc.Service, error) {
	services, err := ss.queries.ListAllServices(ctx)
	if err != nil {
		logger.Error("Failed to list all services", err)
		return nil, err
	}
	return services, nil
}

func (ss *ServiceService) DeleteService(ctx context.Context, id int32) error {
	err := ss.queries.DeleteService(ctx, id)
	if err != nil {
		logger.Error("Failed to delete service", err)
		return err
	}
	return nil
}


func (ss *ServiceService) UpdateServiceFinishStatus(ctx context.Context, id int32, isFinished bool) (sqlc.Service, error) {
	params := sqlc.UpdateServiceFinishStatusParams{
		ID:         id,
		IsFinished: isFinished,
	}

	service, err := ss.queries.UpdateServiceFinishStatus(ctx, params)
	if err != nil {
		logger.Error("Failed to update service finish status", err)
		return sqlc.Service{}, err
	}
	return service, nil
}


func (ss *ServiceService) UpdateServicePaymentStatus(ctx context.Context, id int32, isPaid bool) (sqlc.Service, error) {
	params := sqlc.UpdateServicePaymentStatusParams{
		ID:     id,
		IsPaid: isPaid,
	}

	service, err := ss.queries.UpdateServicePaymentStatus(ctx, params)
	if err != nil {
		logger.Error("Failed to update service payment status", err)
		return sqlc.Service{}, err
	}
	return service, nil
}
