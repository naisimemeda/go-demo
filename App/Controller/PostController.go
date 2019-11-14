package Controller


import (
	"github.com/gin-gonic/gin"
	"go-demo/App/Api"
	"go-demo/App/Models"
	"go-demo/config"
	"net/http"
)

func PostList(ctx *gin.Context) {
	db := config.GetDB()

	var post []Models.Post
	db.Find(&post)
	ctx.JSON(http.StatusOK, Api.Success(post))
	return
}