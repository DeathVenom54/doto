package router

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

var Router *chi.Mux

func init() {
	Router = chi.NewRouter()

	Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello!"))
		if err != nil {
			log.Printf("Error while writing response at /:\n%s", err)
		}
	})
}
