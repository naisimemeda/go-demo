package route

import (
	"github.com/gin-gonic/gin"
	"go-demo/App/Controller"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/user/list", Controller.UserList)
		api.POST("/post/list", Controller.PostList)
		api.POST("/user/create", Controller.CreateUser)
		api.POST("/user/update", Controller.UpdateUser)
		api.POST("/user/delete", Controller.DeleteUser)
	}
	return router
}


