package route

import (
	"github.com/dongwlin/anime-list/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Setup(
	app *fiber.App,
	pingHandler *handler.PingHandler,
	animeHandler *handler.AnimeHandler,
	seasonHandler *handler.SeasonHandler,
	episodeHandler *handler.EpisodeHandler,
	theaterHandler *handler.TheaterHandler,
) {
	app.All("/ping", pingHandler.Ping)

	api := app.Group("/api")

	// anime handlers
	api.Post("/animes", animeHandler.Create)
	api.Get("/animes", animeHandler.List)
	api.Get("/animes/:id", animeHandler.Get)
	api.Put("/animes/:id", animeHandler.Update)
	api.Delete("/animes/:id", animeHandler.Delete)

	// season handlers
	api.Post("/animes/:animeId/seasons", seasonHandler.Create)
	api.Get("/animes/:animeId/seasons", seasonHandler.List)
	api.Get("/seasons/:id", seasonHandler.Get)
	api.Put("/seasons/:id", seasonHandler.Update)
	api.Delete("/seasons/:id", seasonHandler.Delete)

	// episode handlers
	api.Post("/seasons/:seasonId/episodes", episodeHandler.Create)
	api.Get("/seasons/:seasonId/episodes", episodeHandler.List)
	api.Get("/episodes/:id", episodeHandler.Get)
	api.Put("/episodes/:id", episodeHandler.Update)
	api.Delete("/episodes/:id", episodeHandler.Delete)

	// theater handlers
	api.Post("/animes/:animeId/theaters", theaterHandler.Create)
	api.Get("/animes/:animeId/theaters", theaterHandler.List)
	api.Get("/theaters/:id", theaterHandler.Get)
	api.Put("/theaters/:id", theaterHandler.Update)
	api.Delete("/theaters/:id", theaterHandler.Delete)
}
