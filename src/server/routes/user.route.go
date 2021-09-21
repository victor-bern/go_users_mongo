package routes

import (
	"gomongo/src/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(c *gin.Engine) *gin.Engine {
	route := c.Group("api/v1")
	{
		route.Group("users")
		{
			route.GET("All", controllers.GetAll)
			route.POST("create", controllers.Create)
		}
	}

	return c
}
