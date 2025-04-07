package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/josevitorrodriguess/client-manager/internal/config/logger"
	"github.com/josevitorrodriguess/client-manager/internal/jsonutils"
	"go.uber.org/zap"
)

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		logger.Debug("Checking authentication", zap.String("request_id", requestID))

		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			logger.Warn("Unauthorized access attempt", zap.String("request_id", requestID))
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "must be logged in",
			})
			return
		}

		userIDInterface := api.Sessions.Get(r.Context(), "AuthenticatedUserId")
		userID, ok := userIDInterface.(string)
		if !ok {
			logger.Error("Invalid session data", nil, zap.String("request_id", requestID))
			jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"message": "invalid session data",
			})
			return
		}

		logger.Debug("Authentication successful",
			zap.String("request_id", requestID),
			zap.String("user_id", userID))
		next.ServeHTTP(w, r)
	})
}

func (api *Api) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		logger.Debug("Checking admin access", zap.String("request_id", requestID))

		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			logger.Warn("Unauthorized admin access attempt", zap.String("request_id", requestID))
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "must be logged in",
			})
			return
		}

		userIDInterface := api.Sessions.Get(r.Context(), "AuthenticatedUserId")
		userID, ok := userIDInterface.(string)
		if !ok {
			logger.Error("Invalid session data in admin check", nil, zap.String("request_id", requestID))
			jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"message": "invalid session data",
			})
			return
		}

		// Parse the string into UUID
		parsedUserID, err := uuid.Parse(userID)
		if err != nil {
			logger.Error("Invalid user ID format", err, zap.String("request_id", requestID))
			jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"message": "invalid user ID format",
			})
			return
		}

		ok, err = api.UserService.CheckIsAdmin(r.Context(), parsedUserID)
		if err != nil {
			logger.Error("Failed to check admin status", err,
				zap.String("request_id", requestID),
				zap.String("user_id", userID))
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "internal server error",
			})
			return
		}

		if !ok {
			logger.Warn("Non-admin access attempt",
				zap.String("request_id", requestID),
				zap.String("user_id", userID))
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "only admins can access this resource",
			})
			return
		}

		logger.Debug("Admin access granted",
			zap.String("request_id", requestID),
			zap.String("user_id", userID))
		next.ServeHTTP(w, r)
	})
}
