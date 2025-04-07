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

		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "must be logged in",
			})
			return
		}

		userIDInterface := api.Sessions.Get(r.Context(), "AuthenticatedUserId")
		_, ok := userIDInterface.(string)
		if !ok {
			logger.Error("Invalid session data", nil, zap.String("request_id", requestID))
			jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"message": "invalid session data",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *Api) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")

		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
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
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "only admins can access this resource",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
