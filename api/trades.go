package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/lib/pq"
	db "github.com/mstoews/prd-backup-server/db/sqlc"
	// "github.com/mstoews/prd-backup-server/token"
)

// type createAccountRequest struct {
// 	Currency string `json:"currency" binding:"required,currency"`
// }

// func (server *Server) createAccount(ctx *gin.Context) {
// 	var req createAccountRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
// 	arg := db.CreateAccountParams{
// 		Owner:    authPayload.Username,
// 		Currency: req.Currency,
// 		Balance:  0,
// 	}

// 	account, err := server.store.CreateAccount(ctx, arg)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			switch pqErr.Code.Name() {
// 			case "foreign_key_violation", "unique_violation":
// 				ctx.JSON(http.StatusForbidden, errorResponse(err))
// 				return
// 			}
// 		}
// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, account)
// }

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

	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// if account.Owner != authPayload.Username {
	// 	err := errors.New("Trades error")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }

	ctx.JSON(http.StatusOK, trade)
}

type listTradeRequest struct {
	TrdDate  time.Time `form:"trd_date" binding:"required,min=1"`
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTrades(ctx *gin.Context) {
	var req listTradeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	    arg := db.ListTradesParams{
		TrdDate:  req.TrdDate,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	trades, err := server.store.ListTrades(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, trades)
}
