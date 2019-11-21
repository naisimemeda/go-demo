package Controller

import (
	"github.com/gin-gonic/gin"
	"go-demo/App/Api"
	"go-demo/App/Models"
	"go-demo/config"
	"net/http"
	"time"
)

func UserList(ctx *gin.Context) {
	db := config.GetDB()
	var user []Models.UserPost
	db.Preload("Post.Comment").Find(&user)
	ctx.JSON(http.StatusOK, Api.Success(user))
	return
}

func CreateUser(ctx *gin.Context) {
	db := config.GetDB()
	user := Models.User{Name : "xxx", CreatedAt: Api.JSONTime{time.Now()}}
	db.Create(&user)
	ctx.JSON(http.StatusOK, Api.Success("成功"))
	return
}

func UpdateUser(ctx *gin.Context) {
	db := config.GetDB()
	var user Models.User
	db.Where("id = ?", 3).First(&user)
	db.Model(&user).Updates(Models.User{Name:"xxx"})
	ctx.JSON(http.StatusOK, Api.Success(user))
	return
}