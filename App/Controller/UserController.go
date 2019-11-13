package Controller

import (
	"github.com/gin-gonic/gin"
	"go-demo/App/Api"
	"net/http"
)

func UserList(ctx *gin.Context) {
	id := ctx.Query("id")
	ctx.JSON(http.StatusOK, Api.Success(id))
	return
}
