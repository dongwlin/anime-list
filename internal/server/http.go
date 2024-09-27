package server

import (
	"github.com/dongwlin/anime-list/internal/handler"
	"github.com/dongwlin/anime-list/internal/route"
	"github.com/gofiber/fiber/v2"
)

type HttpServer struct {
	app *fiber.App
}

func NewHttpServer(
	pingHandler *handler.PingHandler,
	animeHandler *handler.AnimeHandler,
	seasonHandler *handler.SeasonHandler,
	episodeHandler *handler.EpisodeHandler,
	theaterHandler *handler.TheaterHandler,
) *HttpServer {
	app := fiber.New()
	route.Setup(app, pingHandler, animeHandler, seasonHandler, episodeHandler, theaterHandler)
	return &HttpServer{
		app: app,
	}
}

func (s *HttpServer) Run(addr string) error {
	return s.app.Listen(addr)
}
