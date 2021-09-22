package routes

import (
	"gomongo/src/controllers"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return c
}
