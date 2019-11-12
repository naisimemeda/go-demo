package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, "hello")
			return
		})
	}
	return router
}
