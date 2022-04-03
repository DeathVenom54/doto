package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
	"log"
	"net/http"
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

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("Welcome to the DoTo API!"))
			if err != nil {
				log.Printf("Error while writing response at /:\n%s", err)
			}
		})
	})
}
