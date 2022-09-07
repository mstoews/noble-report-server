package api

import (
	
	"net/http"

	"github.com/gin-gonic/gin"
	// db "github.com/mstoews/prd-backup-server/db/sqlc"
	// "github.com/mstoews/prd-backup-server/token"
)


func (server *Server) listParties(ctx *gin.Context) {
	parties, err := server.store.ListParties(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, parties)
}


