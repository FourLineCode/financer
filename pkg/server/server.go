package server

import (
	"log"

	"github.com/FourLineCode/financer/internal/config"
	"github.com/FourLineCode/financer/pkg/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	db  *gorm.DB
	log log.Logger
}

func New(config *config.Config) *Server {
	app := fiber.New()

	db := initializeDB(config)

	return &Server{app, db, log.Logger{}}
}

func (s *Server) Run(port string) {
	s.initializeRouter()

	s.log.Fatal(s.app.Listen(port))
}

func initializeDB(config *config.Config) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		panic("Failed to connect database!")
	}

	db.AutoMigrate(
		&model.User{},
	)

	return db
}
