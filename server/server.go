package server

import (
	"net/http"

	"github.com/Himany/go-metrics/handlers"
)

func Run(handler *handlers.UpdateHandler) error {
	mux := http.NewServeMux()
	mux.Handle("/update/", handler)
	return http.ListenAndServe(":8080", mux)
}
