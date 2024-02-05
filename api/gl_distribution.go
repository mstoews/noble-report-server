package api

import (
	//"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	"time"
	// db "github.com/mstoews/prd-backup-server/db/sqlc"
	// "github.com/mstoews/prd-backup-server/token"
)


// type getTradeRequest struct {
// 	TrdRecordno int32 `TrdRecordno:"id" binding:"required,min=1"`
// }

// func (server *Server) getDistributionByRef(ctx *gin.Context) {
// 	var req getTradeRequest
// 	if err := ctx.ShouldBindUri(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	trade, err := server.store.getDistributionByRef(ctx, req.TrdRecordno)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}

// 		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, trade)
// }

type DistListResponse struct {
	Account 		int32       `json:"account"`
	Child   		int32       `json:"child"`
	Period  		int32       `json:"period"`
	Description 	string       `json:"description"`
	OpeningBalance 	string       `json:"opening_balance"`
	DebitBalance 	string       `json:"debit_balance"`
	CreditBalance 	string       `json:"credit_balance"`	
	UpdateDate 		time.Time    `json:"update_date"`
	CreateUser  	string       `json:"created_user"`
}


func (server *Server) ListDistLedger(ctx *gin.Context) {
	gldist, err := server.store.ListDistLedger(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gldist)
}


