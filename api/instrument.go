package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// db "github.com/mstoews/prd-backup-server/db/sqlc"
	// "github.com/mstoews/prd-backup-server/token"
)


func (server *Server) ListInstruments(ctx *gin.Context) {
	instruments, err := server.store.ListInstruments(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, instruments)
}

