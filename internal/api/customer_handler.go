package api

import (
	"errors"
	"net/http"

	"github.com/josevitorrodriguess/client-manager/internal/jsonutils"
	"github.com/josevitorrodriguess/client-manager/internal/services"
	"github.com/josevitorrodriguess/client-manager/internal/validators/customer"
)

func (api *Api) HandlerCreatePFCustomer(w http.ResponseWriter, r *http.Request) {
	data, err := jsonutils.DecodeJson[customer.CustomerPFRequest](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
	}

	ok, err := data.IsValid()
	if !ok {
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err)
	}

	id, err := api.CustomerService.CreatePFCustomer(r.Context(), data)
	if err != nil {
		errors.Is(err, services.ErrDuplicatedData)
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"customer_id": id})
}

func (api *Api) HandlerCreatePJCustomer(w http.ResponseWriter, r *http.Request) {
	data, err := jsonutils.DecodeJson[customer.CustomerPJRequest](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ok, err := data.IsValid()
	if !ok {
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := api.CustomerService.CreatePJCustomer(r.Context(), data)
	if err != nil {
		errors.Is(err, services.ErrDuplicatedData)
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"customer_id": id})
}
