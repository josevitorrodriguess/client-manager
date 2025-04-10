package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)
	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.With(api.AdminMiddleware).Post("/register", api.SignUpUserHandler)
				r.Post("/login", api.LoginUserHandler)
				r.With(api.AuthMiddleware).Post("/logout", api.LogoutUserHandler)
			})
			r.Route("/customers", func(r chi.Router) {
				r.With(api.AuthMiddleware, api.AdminMiddleware).Post("/PFcustomer", api.HandlerCreatePFCustomer)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Post("/PJcustomer", api.HandlerCreatePJCustomer)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Post("/addAddress", api.HandlerAddAddressToCostumer)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Get("/getById/{id}", api.HandlerGetCustomerById)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Get("/getAll", api.HandleGetAllCustomers)
			})
		})
	})
}
