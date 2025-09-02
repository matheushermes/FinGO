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
	"github.com/matheushermes/FinGO/internal/models"
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

func GetCryptoMarketChart(crypto, days, interval string) (*models.MarketChart, error) {
	cacheKey := fmt.Sprintf("market_chart:%s:%s:%s", crypto, days, interval)

	if cached, err := cache.Get(cacheKey); err != nil {
		var result models.MarketChart
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return &result, nil
		}
	}

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart?vs_currency=usd&days=%s&interval=%s", crypto, days, interval)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-cg-demo-api-key", configs.API_KEY_COINGECKO)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from CoinGecko: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("CoinGecko API error (%d): %s", resp.StatusCode, string(body))
	}

	var result models.MarketChart
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	cache.Set(cacheKey, string(body), 15*time.Minute)

	return &result, nil
}

func GetCryptoMarketChartRange(crypto string, from, to time.Time) (*models.MarketChart, error) {
	cacheKey := fmt.Sprintf("market_chart_range:%s:%d:%d", crypto, from.Unix(), to.Unix())

	if cached, err := cache.Get(cacheKey); err == nil {
		var result models.MarketChart
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return &result, nil
		}
	}

	url := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart/range?vs_currency=usd&from=%d&to=%d",
	crypto, from.Unix(), to.Unix())

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("x-cg-demo-api-key", configs.API_KEY_COINGECKO)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("coingecko error %d: %s", resp.StatusCode, string(body))
	}

	var result models.MarketChart
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	cache.Set(cacheKey, string(body), 15*time.Minute)

	return &result, nil
}
