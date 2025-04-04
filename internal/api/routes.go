package api

import (
	"github.com/go-chi/chi/v5"
)

func (api *Api) BindRoutes() {
	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.With(api.AdminMiddleware).Post("/register", api.SignUpUserHandler)
				r.Post("/login", api.LoginUserHandler)
				r.With(api.AuthMiddleware).Post("/logout",api.LogoutUserHandler)
			})
		})
	})
}
