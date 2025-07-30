package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/internal/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := auth.ValidateToken(c)

		if err != nil {
			c.AbortWithStatus(401)
		}
		c.Next()
	}
}