package server

import (
	"log"

	"github.com/FourLineCode/financer/config"
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
	s.db = s.initializeDB(config)
	s.Router = s.initializeRouter()

	return s.Router
}

func (s *Server) initializeDB(config *config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		panic("Failed to connect database!")
	}

	// Migrate all models
	db.AutoMigrate(
		&model.Product{},
		&model.User{},
	)

	return db
}

func (s *Server) initializeRouter() *mux.Router {
	r := mux.NewRouter()
	h := handler.New(s.db)

	s.registerRoutes(r, h)

	return r
}
