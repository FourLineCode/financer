package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *Server) initializeRouter() {
	// Middlewares
	s.app.Use(logger.New())
	s.app.Use(cors.New())

	// Routes
	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"msg": "Welcome to Finacner!"})
	})
}
