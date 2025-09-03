package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/internal/auth"
	"github.com/matheushermes/FinGO/internal/database"
	"github.com/matheushermes/FinGO/internal/models"
	"github.com/matheushermes/FinGO/internal/repository"
	"github.com/matheushermes/FinGO/pkg/utils"
)

func CreatePriceCryptoAlert(c *gin.Context) {
	var alert models.PriceAlert

	userId, email, err := auth.ExtractDataFromToken(c)
	if err != nil {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(400, gin.H{
			"error": "failed to create price alert, please check the data provided",
		})
		return
	}

	alert.UserEmail = email

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(500, gin.H{
			"error": "failed to connect to the database",
		})
		return
	}
	defer db.Close()

	repoCrypto := repository.NewCryptosRepository(db)
	crypto, err := repoCrypto.GetCrypto(alert.CryptoID)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "crypto not found in your portfolio",
		})
		return
	}

	if crypto.UserID != userId {
		c.JSON(403, gin.H{
			"error": "you do not have permission to create an alert for this crypto",
		})
		return
	}
	
	alert.Symbol = crypto.Symbol

	targetPrice := utils.CalculateTargetPriceCrypto(alert, crypto.PurchasePriceUSD)
	alert.TargetPrice = targetPrice

	repoAlert := repository.NewAlertsRepository(db)
	alertID, err := repoAlert.CreateAlert(alert)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "error creating price alert: " + err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"id":     alertID,
		"message": "price alert created successfully",
	})

}