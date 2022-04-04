package router

import (
	"github.com/DeathVenom54/doto-backend/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
	"os"
	"time"
)

var Router *chi.Mux

func init() {
	Router = chi.NewRouter()

	if os.Getenv("ENVIRONMENT") == "DEV" {
		Router.Use(middleware.Logger)
	}

	Router.Route("/api/doto", func(r chi.Router) {
		r.Use(httprate.LimitAll(100, time.Minute))

		r.Get("/", handlers.Index)

		r.Route("/auth", authRouter)
	})
}
