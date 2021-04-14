package server

import (
	"net/http"

	"github.com/FourLineCode/financer/src/handler"
)

func (s *Server) IndexRouter(w http.ResponseWriter, r *http.Request) {
	handler.IndexHandler(s.db, w, r)
}

func (s *Server) ProductRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			handler.GetAllProducts(s.db, w, r)
			break
		}
	case http.MethodPost:
		{
			handler.CreateProduct(s.db, w, r)
			break
		}
	}
}

func (s *Server) ProductRouterByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			handler.GetProductByID(s.db, w, r)
			break
		}
	case http.MethodPut:
		{
			handler.UpdateProduct(s.db, w, r)
			break
		}
	case http.MethodDelete:
		{
			handler.DeleteProduct(s.db, w, r)
			break
		}
	}
}
