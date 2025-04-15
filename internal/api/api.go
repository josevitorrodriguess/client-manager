package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/josevitorrodriguess/client-manager/internal/services"
)

type Api struct {
	Router          *chi.Mux
	UserService     services.UserService
	CustomerService services.CustomerService
	ServiceService  services.ServiceService
	Sessions        *scs.SessionManager
}
