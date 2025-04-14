package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
	"github.com/josevitorrodriguess/client-manager/internal/jsonutils"
	"github.com/josevitorrodriguess/client-manager/internal/services"
	"github.com/josevitorrodriguess/client-manager/internal/validators/customer"
	"go.uber.org/zap"
)

func (api *Api) HandlerCreatePFCustomer(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")
	logger.Info("Processing PF customer creation request", zap.String("request_id", requestID))

	data, err := jsonutils.DecodeJson[customer.CustomerPFRequest](r)
	if err != nil {
		logger.Error("Failed to decode PF customer request", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ok, err := data.IsValid()
	if !ok {
		logger.Warn("Invalid PF customer data", zap.String("error", err.Error()), zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := api.CustomerService.CreatePFCustomer(r.Context(), data)
	if err != nil {
		if errors.Is(err, services.ErrDuplicatedData) {
			logger.Warn("Duplicate PF customer data",
				zap.String("email", data.Email),
				zap.String("cpf", data.Cpf),
				zap.String("request_id", requestID))
		} else {
			logger.Error("Failed to create PF customer", err, zap.String("request_id", requestID))
		}
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	logger.Info("PF customer created successfully",
		zap.String("customer_id", id.String()),
		zap.String("name", data.Name),
		zap.String("request_id", requestID))
	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"customer_id": id})
}

func (api *Api) HandlerCreatePJCustomer(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")
	logger.Info("Processing PJ customer creation request", zap.String("request_id", requestID))

	data, err := jsonutils.DecodeJson[customer.CustomerPJRequest](r)
	if err != nil {
		logger.Error("Failed to decode PJ customer request", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ok, err := data.IsValid()
	if !ok {
		logger.Warn("Invalid PJ customer data", zap.String("error", err.Error()), zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := api.CustomerService.CreatePJCustomer(r.Context(), data)
	if err != nil {
		if errors.Is(err, services.ErrDuplicatedData) {
			logger.Warn("Duplicate PJ customer data",
				zap.String("email", data.Email),
				zap.String("cnpj", data.Cnpj),
				zap.String("request_id", requestID))
		} else {
			logger.Error("Failed to create PJ customer", err, zap.String("request_id", requestID))
		}
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	logger.Info("PJ customer created successfully",
		zap.String("customer_id", id.String()),
		zap.String("company_name", data.CompanyName),
		zap.String("request_id", requestID))
	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"customer_id": id})
}

func (api *Api) HandlerAddAddressToCostumer(w http.ResponseWriter, r *http.Request) {
	data, err := jsonutils.DecodeJson[customer.AddAddressRequest](r)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
	}
	id, err := api.CustomerService.AddAddressToCustomer(r.Context(), data)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"address_id": id})
}

func (api *Api) HandlerGetCustomerById(w http.ResponseWriter, r *http.Request) {
	ID := chi.URLParam(r, "id")

	customerID, err := uuid.Parse(ID)
	if err != nil {
		http.Error(w, "invalid customer id", http.StatusBadRequest)
		return
	}

	customer, err := api.CustomerService.GetCustomerDetails(r.Context(), customerID)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{"error": err})
		return
	}
	if customer.ID == uuid.Nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{"message": "no customer found"})
	}
	_ = jsonutils.EncodeJson(w, r, http.StatusOK, customer)
}

func (api *Api) HandleGetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := api.CustomerService.GetAllCustomersDetails(r.Context())
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{"error": err})
	}
	if len(customers) == 0 {
		_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "no customers found"})
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, customers)
}

func (api *Api) HandlerDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id, err := jsonutils.DecodeJson[uuid.UUID](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{"error": err})
	}

	if err = api.CustomerService.DeleteCustomer(r.Context(), id); err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "Customer Deleted Sucessfully "})
	}
}
