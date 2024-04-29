package routes

import (
	"github.com/dimitar-ivanov-93/htmx/handlers"
	"github.com/dimitar-ivanov-93/htmx/middleware"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(router *chi.Mux) {
	router.Get("/signin", handlers.Make(handlers.HandleSigninIndex))
	router.Get("/login", handlers.Make(handlers.HandleSignin))
	router.Get("/callback", handlers.Make(handlers.HandleCallback))

	router.Route("/", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Handle("/*", handlers.Public())
		r.Get("/", handlers.Make(handlers.HandleHome))
	})
}
