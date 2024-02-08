package api

import (
	//"database/sql"
	"net/http"
	"github.com/gin-gonic/gin"
	// db "github.com/mstoews/prd-backup-server/db/sqlc"
	// "github.com/mstoews/prd-backup-server/token"
)

func (server *Server) KBTasks(ctx *gin.Context) {
	kbtask , err := server.store.ListTasks(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, kbtask)
}

func (server *Server) GetTasks(ctx *gin.Context) {
	kbtasklist , err := server.store.GetTaskList(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, kbtasklist)
}

