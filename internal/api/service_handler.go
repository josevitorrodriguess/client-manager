package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
	"github.com/josevitorrodriguess/client-manager/internal/jsonutils"
	"github.com/josevitorrodriguess/client-manager/internal/validators/service"
	"go.uber.org/zap"
)

func (api *Api) HandlerCreateService(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")

	data, err := jsonutils.DecodeJson[service.ServiceRequest](r)
	if err != nil {
		logger.Error("Failed to decode service request", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ok, err := data.IsValid()
	if !ok {
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err)
		return
	}

	serviceID, err := api.ServiceService.CreateService(r.Context(), data)
	if err != nil {
		logger.Error("Failed to create service", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"service_id": serviceID})
}

func (api *Api) HandlerGetServicesByCustomerID(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")
	customerIDStr := chi.URLParam(r,"id")
	if customerIDStr == "" {
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, "customer_id is required")
		return
	}

	customerID, err := uuid.Parse(customerIDStr)
	if err != nil {
		logger.Error("Invalid customer_id format", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, "invalid customer_id format")
		return
	}

	services, err := api.ServiceService.GetServicesByCustomerID(r.Context(), customerID)
	if err != nil {
		logger.Error("Failed to get services for customer", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, services)
}

func (api *Api) HandlerCountServicesByCustomerID(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")
	customerIDStr := chi.URLParam(r, "id")
	if customerIDStr == "" {
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, "customer_id is required")
		return
	}

	customerID, err := uuid.Parse(customerIDStr)
	if err != nil {
		logger.Error("Invalid customer_id format", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, "invalid customer_id format")
		return
	}

	count, err := api.ServiceService.CountServicesByCustomerID(r.Context(), customerID)
	if err != nil {
		logger.Error("Failed to count services for customer", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]int64{"count": count})
}

func (api *Api) HandlerListAllServices(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")

	services, err := api.ServiceService.ListAllServices(r.Context())
	if err != nil {
		logger.Error("Failed to list all services", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, services)
}

func (api *Api) HandlerDeleteService(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, "id is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		logger.Error("Invalid id format", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, "invalid id format")
		return
	}

	err = api.ServiceService.DeleteService(r.Context(), int32(id))
	if err != nil {
		logger.Error("Failed to delete service", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, map[string]string{"message": "service deleted successfully"})
}

func (api *Api) HandlerUpdateServiceFinishStatus(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")

	request, err := jsonutils.DecodeJson[service.UpdateServiceStatusRequest](r)
	if err != nil {
		logger.Error("Failed to decode request", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	service, err := api.ServiceService.UpdateServiceFinishStatus(r.Context(), request.ID, request.Status)
	if err != nil {
		logger.Error("Failed to update service finish status", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, service)
}

func (api *Api) HandlerUpdateServicePaymentStatus(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")

	request, err := jsonutils.DecodeJson[service.UpdateServiceStatusRequest](r)
	if err != nil {
		logger.Error("Failed to decode request", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	service, err := api.ServiceService.UpdateServicePaymentStatus(r.Context(), request.ID, request.Status)
	if err != nil {
		logger.Error("Failed to update service payment status", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusOK, service)
}
