package api

import (
	"github.com/dongwlin/anime-list/internal/api/route"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	app := fiber.New()
	route.Setup(app)
	return &Server{
		app: app,
	}
}

func (s *Server) Run(addr string) error {
	return s.app.Listen(addr)
}
