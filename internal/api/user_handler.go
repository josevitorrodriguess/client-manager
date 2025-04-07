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
	logger.Info("Processing user signup request", zap.String("request_id", requestID))

	data, err := jsonutils.DecodeJson[user.UserRequest](r)
	if err != nil {
		logger.Error("Failed to decode signup request", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ok, err := data.IsValid()
	if !ok {
		logger.Warn("Invalid signup data", zap.String("error", err.Error()), zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err)
		return
	}

	id, err := api.UserService.Create(r.Context(), data)
	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrUsername) {
			logger.Warn("Duplicate email or username", zap.String("email", data.Email), zap.String("request_id", requestID))
		} else {
			logger.Error("Failed to create user", err, zap.String("request_id", requestID))
		}
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
		return
	}

	logger.Info("User created successfully", zap.String("user_id", id.String()), zap.String("request_id", requestID))
	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"user_id": id})
}

func (api *Api) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")
	logger.Info("Processing login request", zap.String("request_id", requestID))

	data, err := jsonutils.DecodeJson[user.UserRequestLogin](r)
	if err != nil {
		logger.Error("Failed to decode login request", err, zap.String("request_id", requestID))
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{"error": "invalid json"})
		return
	}

	id, err := api.UserService.AuthenticateUser(r.Context(), data.Email, string(data.Password))
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			logger.Warn("Invalid login attempt", zap.String("email", data.Email), zap.String("request_id", requestID))
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
	logger.Info("User logged in successfully", zap.String("user_id", id.String()), zap.String("request_id", requestID))
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "logged in sucessfully"})
}

func (api *Api) LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	requestID := r.Header.Get("X-Request-ID")
	logger.Info("Processing logout request", zap.String("request_id", requestID))

	err := api.Sessions.RenewToken(r.Context())
	if err != nil {
		logger.Error("Failed to renew session token during logout", err, zap.String("request_id", requestID))
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	logger.Info("User logged out successfully", zap.String("request_id", requestID))
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "logged out sucessfully"})
}
