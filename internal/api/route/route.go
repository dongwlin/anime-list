package route

import (
	"github.com/dongwlin/anime-list/internal/api/handler"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.All("/ping", handler.Ping)

	api := app.Group("/api")

	// anime handlers
	api.Post("/animes", handler.CreateAnime)
	api.Get("/animes", handler.ListAnime)
	api.Get("/animes/:id", handler.GetAnime)
	api.Put("/animes/:id", handler.UpdateAnime)
	api.Delete("/animes/:id", handler.DeleteAnime)

	// season handlers
	api.Post("/animes/:animeId/seasons", handler.CreateSeason)
	api.Get("/animes/:animeId/seasons", handler.ListSeason)
	api.Get("/seasons/:id", handler.GetSeason)
	api.Put("/seasons/:id", handler.UpdateSeason)
	api.Delete("/seasons/:id", handler.DeleteSeason)

	// episode handlers
	api.Post("/seasons/:seasonId/episodes", handler.CreateEpisode)
	api.Get("/seasons/:seasonId/episodes", handler.ListEpisode)
	api.Get("/episodes/:id", handler.GetEpisode)
	api.Put("/episodes/:id", handler.UpdateEpisode)
	api.Delete("/episodes/:id", handler.DeleteEpisode)

	// theater handlers
	api.Post("/animes/:animeId/theaters", handler.CreateTheater)
	api.Get("/animes/:animeId/theaters", handler.ListTheater)
	api.Get("/theaters/:id", handler.GetTheater)
	api.Put("/theaters/:id", handler.UpdateTheater)
	api.Delete("/theaters/:id", handler.DeleteTheater)
}
