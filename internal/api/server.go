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

	api.POST("/animes", server.createAnime)
	api.GET("animes", server.listAnime)
	api.GET("/animes/:id", server.getAnime)
	api.PUT("/animes/:id", server.updateAnime)
	api.DELETE("/animes/:id", server.deleteAnime)

	server.router = router
	return server
}

// Start runs the HTTP Server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
