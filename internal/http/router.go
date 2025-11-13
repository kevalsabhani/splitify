package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevalsabhani/splitify/internal/http/handler"
)

func NewRouter() http.Handler {
	r := mux.NewRouter()

	healthHandler := handler.NewHealthHandler()
	fmt.Println("here")
	r.HandleFunc("/healthcheck", healthHandler.Healthcheck)

	// v1Router := r.PathPrefix("/api/v1").Subrouter()

	return r
}
