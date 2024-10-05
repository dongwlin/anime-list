package route

import (
	"github.com/dongwlin/anime-list/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Setup(
	app *fiber.App,
	handler *handler.Handler,
) {
	app.All("/ping", handler.Ping.Ping)

	api := app.Group("/api")

	// anime handlers
	api.Post("/animes", handler.Anime.Create)
	api.Get("/animes", handler.Anime.List)
	api.Get("/animes/search", handler.Anime.Search)
	api.Get("/animes/:id", handler.Anime.Get)
	api.Put("/animes/:id", handler.Anime.Update)
	api.Delete("/animes/:id", handler.Anime.Delete)

	// season handlers
	api.Post("/animes/:animeId/seasons", handler.Season.Create)
	api.Get("/animes/:animeId/seasons", handler.Season.List)
	api.Get("/animes/:animeId/seasons/search", handler.Season.Search)
	api.Get("/seasons/:id", handler.Season.Get)
	api.Put("/seasons/:id", handler.Season.Update)
	api.Delete("/seasons/:id", handler.Season.Delete)

	// episode handlers
	api.Post("/seasons/:seasonId/episodes", handler.Episode.Create)
	api.Get("/seasons/:seasonId/episodes", handler.Episode.List)
	api.Get("/seasons/:seasonId/episodes/search", handler.Episode.Search)
	api.Get("/episodes/:id", handler.Episode.Get)
	api.Put("/episodes/:id", handler.Episode.Update)
	api.Delete("/episodes/:id", handler.Episode.Delete)

	// theater handlers
	api.Post("/animes/:animeId/theaters", handler.Theater.Create)
	api.Get("/animes/:animeId/theaters", handler.Theater.List)
	api.Get("/animes/:animeId/theaters/search", handler.Theater.Search)
	api.Get("/theaters/:id", handler.Theater.Get)
	api.Put("/theaters/:id", handler.Theater.Update)
	api.Delete("/theaters/:id", handler.Theater.Delete)
}
