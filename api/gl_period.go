package api

import (
	//"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	// db "github.com/mstoews/prd-backup-server/db/sqlc"
	// "github.com/mstoews/prd-backup-server/token"
)

func (server *Server) GlPeriod(ctx *gin.Context) {
	glPeriod , err := server.store.ListPeriods(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, glPeriod)
}
