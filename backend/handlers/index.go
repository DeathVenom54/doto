package handlers

import (
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Welcome to the DoTo API!"))
	if err != nil {
		log.Printf("Error while writing response at /\n%s", err)
	}
}
