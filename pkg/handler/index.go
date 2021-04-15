package handler

import (
	"net/http"
)

type Welcome struct {
	Message string `json:"message"`
}

func (h *Handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, http.StatusOK, Welcome{Message: "Welcome to our API!"})
}

type Success struct {
	Success bool `json:"success"`
}

func (h *Handler) ApiIndexHandler(w http.ResponseWriter, r *http.Request) {
	ResponseJSON(w, http.StatusOK, Success{Success: true})
}
