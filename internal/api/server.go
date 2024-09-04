package api

import (
	"github.com/dongwlin/anime-list/internal/db"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for anime-list service.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) *Server {
	server := &Server{store: store}

	router := gin.Default()

	router.Any("/ping", server.ping)

	api := router.Group("/api")
	// animes
	api.POST("/animes", server.createAnime)
	api.GET("animes", server.listAnime)
	api.GET("/animes/:animeId", server.getAnime)
	api.PUT("/animes/:animeId", server.updateAnime)
	api.DELETE("/animes/:animeId", server.deleteAnime)
	// seasons
	api.POST("/animes/:animeId/seasons", server.createSeason)
	api.GET("/animes/:animeId/seasons", server.listSeason)
	api.GET("/seasons/:seasonId", server.getSeason)
	api.PUT("/seasons/:seasonId", server.updateSeason)
	api.DELETE("/seasons/:seasonId", server.deleteSeason)
	// episode
	api.POST("/seasons/:seasonId/episodes", server.createEpisode)
	api.GET("/seasons/:seasonId/episodes", server.listEpisodes)
	api.GET("/episodes/:episodeId", server.getEpisode)
	api.PUT("/episodes/:episodeId", server.updateEpisode)
	api.DELETE("/episodes/:episodeId", server.deleteEpisode)
	// theater
	api.POST("/animes/:animeId/theaters", server.createTheater)
	api.GET("/animes/:animeId/theaters", server.listTheater)
	api.GET("/theaters/:theaterId", server.getTheater)
	api.PUT("/theaters/:theaterId", server.updateTheater)
	api.DELETE("/theaters/:theaterId", server.deleteTheater)

	server.router = router
	return server
}

// Start runs the HTTP Server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
