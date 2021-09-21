package routes

import "github.com/gin-gonic/gin"

func ConfigRoutes(c *gin.Engine) *gin.Engine {
	route := c.Group("api/v1")
	{
		route.Group("users")
		{
			route.POST("create")
		}
	}

	return c
}
