package api

import (
	//"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	// db "github.com/mstoews/prd-backup-server/db/sqlc"
	// "github.com/mstoews/prd-backup-server/token"
)

func (server *Server) GlJournalDetail(ctx *gin.Context) {
	gljournalDetails , err := server.store.ListJournalDetails(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gljournalDetails)
}


func (server *Server) GlJournalHeader(ctx *gin.Context) {
	glJournalHeader , err := server.store.ListJournalHeader(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, glJournalHeader)
}


