package Controller

import (
	"github.com/gin-gonic/gin"
	"go-demo/App/Api"
	"go-demo/config"
	"net/http"
)

type User struct {
	ID        uint `gorm:"primary_key" json:"id"`
	Name      string `json:"name"`
	CreatedAt Api.JSONTime `json:"created_at"`
}

func UserList(ctx *gin.Context) {
	db := config.GetDB()
	var user []User
	db.Find(&user)
	ctx.JSON(http.StatusOK, Api.Success(user))
	return
}