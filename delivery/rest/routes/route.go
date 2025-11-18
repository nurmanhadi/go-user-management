package routes

import (
	"user-management/delivery/rest/handler"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	Router      *chi.Mux
	UserHandler *handler.UserHandler
}

func (r *Router) New() {
	r.Router.Route("/api", func(api chi.Router) {
		api.Route("/users", func(users chi.Router) {
			users.Put("/{id}", r.UserHandler.UserChangeProfile)
			users.Get("/{username}", r.UserHandler.UserGetByUsername)

			users.Route("/services", func(services chi.Router) {
				services.Get("/ids", r.UserHandler.UserGetBySliceId)
				services.Get("/count/{id}", r.UserHandler.UserCountById)
				services.Get("/{id}", r.UserHandler.UserGetById)
			})
		})
	})
}
