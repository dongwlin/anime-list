package server

import (
	"github.com/dongwlin/anime-list/internal/handler"
	"github.com/dongwlin/anime-list/internal/middleware"
	"github.com/dongwlin/anime-list/internal/route"
	"github.com/gofiber/fiber/v2"
)

type HttpServer struct {
	app *fiber.App
}

func NewHttpServer(
	handler *handler.Handler,
) *HttpServer {
	app := fiber.New()
	middleware.Setup(app)
	route.Setup(app, handler)
	return &HttpServer{
		app: app,
	}
}

func (s *HttpServer) Run(addr string) error {
	return s.app.Listen(addr)
}
