package server

import (
	"log"
	"net/http"

	"github.com/FourLineCode/financer/pkg/config"
	"github.com/FourLineCode/financer/pkg/handler"
	"github.com/FourLineCode/financer/pkg/model"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	Router *mux.Router
}

func (s *Server) Initialize(config *config.Config) *mux.Router {
	s.db = s.InitializeDB(config)
	s.Router = s.InitializeRouter()

	return s.Router
}

func (s *Server) InitializeDB(config *config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		panic("Failed to connect database!")
	}

	db.AutoMigrate(&model.Product{})
	return db
}

func (s *Server) InitializeRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/", s.IndexRouter).Methods(http.MethodGet)

	router.HandleFunc("/api/product", s.ProductRouter).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/api/product/{id}", s.ProductRouterByID).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	router.NotFoundHandler = http.HandlerFunc((handler.NotFoundHandler))

	return router
}
