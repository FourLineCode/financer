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

	// Migrate all models
	db.AutoMigrate(
		&model.Product{},
		&model.User{},
	)

	return db
}

func (s *Server) InitializeRouter() *mux.Router {
	r := mux.NewRouter()
	h := handler.New(s.db)

	s.RegisterRoutes(r, h)

	return r
}
