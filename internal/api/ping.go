package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong!!!!!")
}
