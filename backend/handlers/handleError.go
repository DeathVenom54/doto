package handlers

import (
	logger "github.com/sirupsen/logrus"
	"net/http"
)

func handleError(err error, w http.ResponseWriter, code int) {
	logger.Errorln(err)
	w.WriteHeader(code)

	_, writeErr := w.Write([]byte(err.Error()))
	if writeErr != nil {
		logger.Errorln(writeErr)
	}
}
