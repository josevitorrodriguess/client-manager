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

}
