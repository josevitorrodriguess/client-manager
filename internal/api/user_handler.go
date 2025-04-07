package api

import (
	"errors"
	"net/http"

	"github.com/josevitorrodriguess/client-manager/internal/jsonutils"
	"github.com/josevitorrodriguess/client-manager/internal/services"
	"github.com/josevitorrodriguess/client-manager/internal/validators/user"
)

func (api *Api) SignUpUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := jsonutils.DecodeJson[user.UserRequest](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err.Error())
	}

	ok, err := data.IsValid()
	if !ok {
		_ = jsonutils.EncodeJson(w, r, http.StatusBadRequest, err)
	}

	id, err := api.UserService.Create(r.Context(), data)
	if err != nil {
		errors.Is(err, services.ErrDuplicatedEmailOrUsername)
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, err)
	}
	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{"user_id": id})
}

func (api *Api) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	data, err := jsonutils.DecodeJson[user.UserRequestLogin](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{"error": "invalid json"})
		return
	}

	id, err := api.UserService.AuthenticateUser(r.Context(), data.Email, string(data.Password))
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{"error": "invalid credentials"})
			return
		}
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	err = api.Sessions.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	api.Sessions.Put(r.Context(), "AuthenticatedUserId", id.String())
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "logged in sucessfully"})

}

func (api *Api) LogoutUserHandler(w http.ResponseWriter, r *http.Request) {
	err := api.Sessions.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{"error": "internal server error"})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{"message": "logged out sucessfully"})
}
