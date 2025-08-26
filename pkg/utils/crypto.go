package utils

import (
	"fmt"

	"github.com/matheushermes/FinGO/internal/models"
)

func EnrichCryptoWithPrice(crypto *models.Crypto) error {
	price, err := GetPriceFromCoinGecko(crypto.Symbol)
	if err != nil {
		return fmt.Errorf("failed to fetch current price for %s: %w", crypto.Symbol, err)
	}

	variation := ((price - crypto.PurchasePriceUSD) / crypto.PurchasePriceUSD) * 100
	currentTotalValue := price * crypto.Amount

	crypto.CurrentPriceUSD = price
	crypto.VariationPercent = variation
	crypto.CurrentTotalValueUSD = currentTotalValue

	return nil
}
