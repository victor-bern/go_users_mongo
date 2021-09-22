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
			route.POST("login", controllers.Login)
			route.GET("all", controllers.GetAll)
			route.POST("create", controllers.Create)
			route.PUT("addOrder/:id", controllers.AddOrder)
		}
	}

	return c
}
