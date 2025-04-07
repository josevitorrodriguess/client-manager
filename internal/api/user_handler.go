package api

import (
	"errors"
	"net/http"

	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
	"github.com/josevitorrodriguess/client-manager/internal/jsonutils"
	"github.com/josevitorrodriguess/client-manager/internal/services"
	"github.com/josevitorrodriguess/client-manager/internal/validators/user"
	"go.uber.org/zap"
)

func (api *Api) SignUpUserHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")

	data, err := jsonutils.DecodeJson[user.UserRequest](r)
	if err != nil {
		logger.Error("Failed to decode signup request", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ok, err := data.IsValid()
	if !ok {
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := api.UserService.Create(r.Context(), data)
	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrUsername) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		logger.Error("Failed to create user", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"user_id": id})
}

func (api *Api) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")

	data, err := jsonutils.DecodeJson[user.UserRequestLogin](r)
	if err != nil {
		logger.Error("Failed to decode login request", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{"error": "invalid json"})
		return
	}

	id, err := api.UserService.AuthenticateUser(r.Context(), data.Email, string(data.Password))
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{"error": "invalid credentials"})
			return
		}
		logger.Error("Authentication error", err, zap.String("request_id", requestID))
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	err = api.Sessions.RenewToken(r.Context())
	if err != nil {
		logger.Error("Failed to renew session token", err, zap.String("request_id", requestID))
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	api.Sessions.Put(r.Context(), "AuthenticatedUserId", id.String())
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "logged in sucessfully"})
}

func (api *Api) LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")

	err := api.Sessions.RenewToken(r.Context())
	if err != nil {
		logger.Error("Failed to renew session token during logout", err, zap.String("request_id", requestID))
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "logged out sucessfully"})
}
