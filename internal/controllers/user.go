package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/internal/database"
	"github.com/matheushermes/FinGO/internal/models"
	"github.com/matheushermes/FinGO/internal/repository"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(422, gin.H{
			"error": "failed to register user, please check the data provided",
		})
		return
	}

	if err := user.IsValid("register"); err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to connect to the database",
		})
		return
	}
	defer db.Close()

	repo := repository.NewUsersRepository(db)
	userID, err := repo.Create(user)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to create user in the database",
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "user registered successfully",
		"userID":  userID,
	})
}