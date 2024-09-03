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

	router.GET("/ping", server.ping)

	server.router = router
	return server
}

// Start runs the HTTP Server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
