package utils

import "github.com/matheushermes/FinGO/internal/models"

func CalculateTargetPriceCrypto(alert models.PriceAlert, purchasePrice float64) float64 {
	if alert.Direction == "above" {
		return purchasePrice * (1 + alert.PercentageChange/100)
	}
	//below
	return purchasePrice * (1 - alert.PercentageChange/100)
}
