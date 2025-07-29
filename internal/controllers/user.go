package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/internal/models"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(422, gin.H{
			"error": "failed to register user, please check the data provided",
		})
		return
	}

	
}