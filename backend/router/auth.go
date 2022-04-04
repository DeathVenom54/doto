package router

import (
	"github.com/DeathVenom54/doto-backend/handlers"
	"github.com/go-chi/chi"
)

func authRouter(r chi.Router) {
	r.Post("/signup", handlers.Signup)
}
