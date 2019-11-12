package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserList(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, "xxxx")
	return

}