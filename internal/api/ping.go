package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong!!!!!")
}
