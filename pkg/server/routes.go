package server

import (
	"net/http"

	"github.com/FourLineCode/financer/pkg/handler"
	"github.com/FourLineCode/financer/pkg/middleware"
	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes(r *mux.Router, h *handler.Handler) {
	// Index Routes
	r.HandleFunc("/", h.IndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/api", h.ApiIndexHandler).Methods(http.MethodGet)

	// Product Routes
	r.HandleFunc("/api/product", middleware.Authenticate(h.GetAllProducts)).Methods(http.MethodGet)
	r.HandleFunc("/api/product", middleware.Authenticate(h.CreateProduct)).Methods(http.MethodPost)
	r.HandleFunc("/api/product/{id}", middleware.Authenticate(h.GetProductByID)).Methods(http.MethodGet)
	r.HandleFunc("/api/product/{id}", middleware.Authenticate(h.UpdateProduct)).Methods(http.MethodPut)
	r.HandleFunc("/api/product/{id}", middleware.Authenticate(h.DeleteProduct)).Methods(http.MethodDelete)

	r.NotFoundHandler = http.HandlerFunc((h.NotFoundHandler))
}
