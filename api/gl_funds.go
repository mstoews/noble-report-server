package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (server *Server) ListFunds(ctx *gin.Context) {
	glfunds , err := server.store.ListFunds(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, glfunds)
}


