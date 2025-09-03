package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/internal/controllers"
	"github.com/matheushermes/FinGO/internal/server/middlewares"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("register", controllers.RegisterUser)
			user.POST("login", controllers.Login)
		}

		crypto := main.Group("crypto", middlewares.AuthMiddleware())
		{
			crypto.POST("actives", controllers.RegisterActives)
			crypto.GET("actives", controllers.GetCryptos)
			crypto.GET("actives/:id", controllers.GetCrypto)
			crypto.GET(":crypto/history", controllers.GetCryptoHistory)
			crypto.GET(":crypto/history/range", controllers.GetCryptoHistoryRange)
		}

		alert := main.Group("alert", middlewares.AuthMiddleware())
		{
			alert.POST("crypto", controllers.CreatePriceCryptoAlert)
		}
	}

	return router
}