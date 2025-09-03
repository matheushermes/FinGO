package controllers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/matheushermes/FinGO/internal/auth"
	"github.com/matheushermes/FinGO/internal/database"
	"github.com/matheushermes/FinGO/internal/models"
	"github.com/matheushermes/FinGO/internal/repository"
	"github.com/matheushermes/FinGO/pkg/utils"
)

func RegisterActives(c *gin.Context) {
	var crypto models.Crypto

	userID, _, err := auth.ExtractDataFromToken(c)
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
	userID, _, err := auth.ExtractDataFromToken(c)
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

	var failedUpdates []string

	for i := range cryptos {
		if err = utils.EnrichCryptoWithPrice(&cryptos[i]); err != nil {
			failedUpdates = append(failedUpdates, fmt.Sprintf("%s: %v", cryptos[i].Symbol, err))
			continue
		}

		if err := repo.UpdateCrypto(cryptos[i]); err != nil {
			failedUpdates = append(failedUpdates, fmt.Sprintf("DB update failed for %s: %v", cryptos[i].Symbol, err))
			continue
		}
	}

	if len(failedUpdates) > 0 {
		for _, msg := range failedUpdates {
			log.Printf("[GetCryptos] update failed: %s", msg)
		}
	}

	c.JSON(200, gin.H{
		"cryptos": cryptos,
	})
	
}

func GetCrypto(c *gin.Context) {
	parameter := c.Param("id")
	idCrypto, err := strconv.ParseUint(parameter, 10, 64)
	if err != nil {
		c.JSON(404, gin.H {
			"error": "invalid crypto ID",
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
	crypto, err := repo.GetCrypto(idCrypto)
	if err != nil {
		c.JSON(500, gin.H {
			"error": err.Error(),
		})
		return
	}

	if err = utils.EnrichCryptoWithPrice(&crypto); err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("failed to enrich crypto with price: %v", err),
		})
		return
	}

	c.JSON(200, gin.H{
		"crypto": crypto,
	})
}

func GetCryptoHistory(c *gin.Context) {
	crypto := c.Param("crypto")
	days := c.DefaultQuery("days", "30")
	interval := c.DefaultQuery("interval", "daily")

	data, err := utils.GetCryptoMarketChart(crypto, days, interval)
	if err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("failed to get crypto market chart: %v", err),
		})
		return
	}

	c.JSON(200, gin.H{
		"crypto": crypto,
		"days": days,
		"interval": interval,
		"history": data,
	})
}

func GetCryptoHistoryRange(c *gin.Context) {
	crypto := c.Param("crypto")
	fromStr := c.Query("from")
	toStr := c.Query("to")

	from, err := time.Parse("2006-01-02", fromStr)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "invalid from date, expected format YYYY-MM-DD",
		})
		return
	}

	to, err := time.Parse("2006-01-02", toStr)
	if err != nil {
		c.JSON(404, gin.H{
			"error": "invalid to date, expected format YYYY-MM-DD",
		})
		return
	}

	data, err := utils.GetCryptoMarketChartRange(crypto, from, to)
	if err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Sprintf("failed to fetch history range: %v", err),
		})
		return
	}

	
	c.JSON(200, gin.H{
		"crypto":  crypto,
		"from":    fromStr,
		"to":      toStr,
		"history": data,
	})
}