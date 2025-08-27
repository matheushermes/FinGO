package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/matheushermes/FinGO/configs"
	"github.com/matheushermes/FinGO/internal/cache"
)

func GetPriceFromCoinGecko(symbol string) (float64, error) {
	symbol = strings.ToLower(symbol)
	cacheKey := "price:" + symbol

	if val, err := cache.Get(cacheKey); err == nil {
		var price float64
		fmt.Sscanf(val, "%f", &price)
		return price, nil
	}

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?symbols=%s&vs_currencies=usd", symbol)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create request: %w", err)
	}

	// Insira sua API key aqui
	req.Header.Set("x-cg-demo-api-key", configs.API_KEY_COINGECKO)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch data from CoinGecko: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("CoinGecko API error (%d): %s", resp.StatusCode, string(bodyBytes))
	}

	var result map[string]map[string]float64
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}

	price, ok := result[symbol]["usd"]
	if !ok {
		return 0, fmt.Errorf("symbol %s not found in CoinGecko response", symbol)
	}

	cache.Set(cacheKey, fmt.Sprintf("%f", price), 15*time.Minute)
	
	return price, nil
}
