package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type getPeriodNo struct {
	Ref int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) ListDistributionLedgerByPeriod(ctx *gin.Context) {
	var req getPeriodNo
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	gldist, err := server.store.ListDistLedgerByPeriod(ctx, req.Ref)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gldist)
}
