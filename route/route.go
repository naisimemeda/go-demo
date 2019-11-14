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
	}
	return router
}


