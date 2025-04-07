package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/josevitorrodriguess/client-manager/internal/jsonutils"
)

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "must be logged in",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *Api) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "must be logged in",
			})
			return
		}

		userIDInterface := api.Sessions.Get(r.Context(), "AuthenticatedUserId")
		userID, ok := userIDInterface.(string)
		if !ok {
			jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"message": "invalid session data",
			})
			return
		}

		// Parse the string into UUID
		parsedUserID, err := uuid.Parse(userID)
		if err != nil {
			jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
				"message": "invalid user ID format",
			})
			return
		}

		ok, err = api.UserService.CheckIsAdmin(r.Context(), parsedUserID)
		if err != nil {
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
