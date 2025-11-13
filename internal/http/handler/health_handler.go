package handler

import (
	"net/http"

	"github.com/kevalsabhani/splitify/pkg/response"
)

type healthHandler struct{}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Healthcheck(w http.ResponseWriter, r *http.Request) {
	response.WriteJSON(w, http.StatusOK, map[string]string{"health": "OK"}, nil)
}
