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
				r.With(api.AuthMiddleware, api.AdminMiddleware).Post("/pf", api.HandlerCreatePFCustomer)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Post("/pj", api.HandlerCreatePJCustomer)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Post("/address", api.HandlerAddAddressToCostumer)
				r.Get("/{id}", api.HandlerGetCustomerById)
				r.Get("/", api.HandleGetAllCustomers)
			})

			r.Route("/services", func(r chi.Router) {
				r.With(api.AuthMiddleware, api.AdminMiddleware).Post("/", api.HandlerCreateService)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Get("/", api.HandlerListAllServices)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Get("/customer/{id}", api.HandlerGetServicesByCustomerID)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Get("/count/{id}", api.HandlerCountServicesByCustomerID)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Delete("/", api.HandlerDeleteService)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Patch("/finish", api.HandlerUpdateServiceFinishStatus)
				r.With(api.AuthMiddleware, api.AdminMiddleware).Patch("/payment", api.HandlerUpdateServicePaymentStatus)
			})
		})
	})
}
