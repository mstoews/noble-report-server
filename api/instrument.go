package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//  db "github.com/mstoews/prd-backup-server/db/sqlc"
	//  "github.com/mstoews/prd-backup-server/token"
)


func (server *Server) ListInstruments(ctx *gin.Context) {
	instruments, err := server.store.ListInstruments(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, instruments)
}


type getInstrumentRequest struct {
	Ref string `uri:"id" binding:"required,min=1"`
}

func (server *Server) GetInstrumentByRef(ctx *gin.Context) {

	var req getInstrumentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	instrument, err := server.store.GetInstrumentsByRef(ctx, req.Ref)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, instrument)
}


