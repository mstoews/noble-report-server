package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (server *Server) GLAccountTypes(ctx *gin.Context) {
	gl_account_types , err := server.store.ListAccountTypes(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gl_account_types)
}


