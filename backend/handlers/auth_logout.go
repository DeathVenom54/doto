package handlers

import (
	logger "github.com/sirupsen/logrus"
	"net/http"
)

func Logout(w http.ResponseWriter, _ *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "user", Value: "", HttpOnly: true, MaxAge: -1})

	_, err := w.Write([]byte("success"))
	if err != nil {
		logger.Errorf("Error while writing response at /auth/logout\n%s\n", err)
	}
}
