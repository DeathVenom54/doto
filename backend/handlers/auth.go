package handlers

import (
	"log"
	"net/http"
)

func Signup(w http.ResponseWriter, _ *http.Request) {
	// TODO implement this
	_, err := w.Write([]byte("hee hee"))
	if err != nil {
		log.Printf("Error while writing response at /auth/signup\n%s", err)
	}
}
