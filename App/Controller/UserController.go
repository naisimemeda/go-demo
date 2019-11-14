package Controller

import (
	"github.com/gin-gonic/gin"
	"go-demo/App/Api"
	"go-demo/App/Models"
	"go-demo/config"
	"net/http"
)

func UserList(ctx *gin.Context) {
	db := config.GetDB()
	var user []Models.User
	db.Preload("Post").Preload("Comment").Find(&user)
	ctx.JSON(http.StatusOK, Api.Success(user))
	return
}