package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/internal/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("register", controllers.RegisterUser)
			user.POST("login", controllers.Login)
		}

		crypto := main.Group("crypto")
		{
			crypto.POST("actives", controllers.RegisterActives)
		}
	}

	return router
}