package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	// db "github.com/mstoews/prd-backup-server/db/sqlc"
	// "github.com/mstoews/prd-backup-server/token"
)

type CreateTradeParams struct {
	TrdRecordno         int32     `json:"trd_recordno"`
	TrdGlosstraderef    int32     `json:"trd_glosstraderef"`
	TrdVersiono         int32     `json:"trd_versiono"`
	TrdOrigin           string    `json:"trd_origin"`
	TrdTradetype        string    `json:"trd_tradetype"`
	TrdSettlementstatus string    `json:"trd_settlementstatus"`
	TrdTradestatus      string    `json:"trd_tradestatus"`
	TrdOriginversion    int32     `json:"trd_originversion"`
	TrdDate             string    `json:"trd_date"`
}

type getTradeRequest struct {
	TrdRecordno int32 `TrdRecordno:"id" binding:"required,min=1"`
}

func (server *Server) getTrade(ctx *gin.Context) {
	var req getTradeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	trade, err := server.store.GetTrade(ctx, req.TrdRecordno)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, trade)
}

func (server *Server) listAllTrades(ctx *gin.Context) {
	trades, err := server.store.ListAllTrades(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, trades)
}


