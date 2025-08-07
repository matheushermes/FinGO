package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/internal/auth"
	"github.com/matheushermes/FinGO/internal/database"
	"github.com/matheushermes/FinGO/internal/models"
	"github.com/matheushermes/FinGO/internal/repository"
	"github.com/matheushermes/FinGO/internal/utils"
)

func RegisterActives(c *gin.Context) {
	var crypto models.Crypto

	userID, err := auth.ExtractDataFromToken(c)
	if err != nil {
		c.JSON(401, gin.H {
			"error": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&crypto); err != nil {
		c.JSON(400, gin.H {
			"error": "failed to register crypto, please check the data provided",
		})
		return
	}

	crypto.UserID = userID

	if err := crypto.ValidationsCryptos(); err != nil {
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

	repo := repository.NewCryptosRepository(db)
	cryptoID, err := repo.Create(crypto)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "crypto registered successfully",
		"cryptoID": cryptoID,
	})
}

func GetCryptos(c *gin.Context) {
	userID, err := auth.ExtractDataFromToken(c)
	if err != nil {
		c.JSON(401, gin.H {
			"error": err.Error(),
		})
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(500, gin.H {
			"error": "failed to connect to the database",
		})
		return
	}
	defer db.Close()

	repo := repository.NewCryptosRepository(db)
	cryptos, err := repo.GetAllCryptos(userID)
	if err != nil {
		c.JSON(500, gin.H {
			"error": err.Error(),
		})
		return
	}

	for i := range cryptos {
		currentPrice, err := utils.GetPrinceFromCoinGecko(cryptos[i].Symbol)
		if err != nil {
			c.JSON(500, gin.H{
				"error": fmt.Sprintf("failed to fetch current price for %s: %v", cryptos[i].Symbol, err),
			})
			return
		}

		variation := ((currentPrice - cryptos[i].PurchasePriceUSD) / cryptos[i].PurchasePriceUSD) * 100
		currentTotalValue := currentPrice * cryptos[i].Amount

		cryptos[i].CurrentPriceUSD = currentPrice
		cryptos[i].VariationPercent = variation
		cryptos[i].CurrentTotalValueUSD = currentTotalValue

		fmt.Println(cryptos[i])
	}

	c.JSON(200, gin.H{
		"cryptos": cryptos,
	})
	
}