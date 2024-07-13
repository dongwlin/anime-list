package api

import "github.com/gin-gonic/gin"

// Server serves HTTP requests for anime-list service.
type Server struct {
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer() *Server {
	server := &Server{}

	router := gin.Default()

	router.GET("/ping", server.ping)

	server.router = router
	return server
}

// Start runs the HTTP Server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
