package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/internal/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		crypto := main.Group("crypto")
		{
			crypto.POST("actives", controllers.RegisterActives)
		}
	}

	return router
}