package handler

import "net/http"

func (h *Handler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ResponseError(w, http.StatusNotFound, "Route not found")
}
