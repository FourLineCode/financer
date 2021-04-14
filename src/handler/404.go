package handler

import "net/http"

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ResponseError(w, http.StatusNotFound, "Route not found")
}
