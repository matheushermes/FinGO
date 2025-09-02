package models

import (
	"errors"
	"strings"
	"time"
)

type MarketChart struct {
	Prices       [][]float64 `json:"prices"`
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

type Crypto struct {
	ID uint64 					`json:"id,omitempty"`
	UserID uint64 				`json:"userID,omitempty"`
	Name string 				`json:"name,omitempty"`
	Symbol string 				`json:"symbol,omitempty"`
	Amount float64 				`json:"amount,omitempty"`
	PurchasePriceUSD float64 	`json:"purchase_price_usd,omitempty"`
	VariationPercent float64 	`json:"variation_percent"`
	CurrentPriceUSD float64 	`json:"current_price_usd"`
	CurrentTotalValueUSD float64 `json:"current_total_value_usd"`
	PurchaseDate time.Time 		`json:"purchase_date"`
	IsSolid bool 				`json:"is_solid,omitempty"`
	Notes string 				`json:"notes,omitempty"`
	CreatedAt time.Time 		`json:"created_at,omitempty"`
	UpdatedAt time.Time 		`json:"updated_at,omitempty"`
}

func (c *Crypto) ValidationsCryptos() error {
	c.trimSpaces()

	if err := c.validateFields(); err != nil {
		return err
	}

	return nil
}

func (c *Crypto) trimSpaces() {
	c.Name = strings.TrimSpace(c.Name)
	c.Symbol = strings.TrimSpace(c.Symbol)
	c.Notes = strings.TrimSpace(c.Notes)
}

func (c *Crypto) validateFields() error {
	if c.Name == "" {
		return errors.New("name cannot be blank")
	}
	if c.Symbol == "" {
		return errors.New("symbol cannot be blank")
	}
	if c.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if c.PurchasePriceUSD < 0 {
		return errors.New("purchase price must be non-negative")
	}
	if c.PurchaseDate.IsZero() {
		return errors.New("purchase date cannot be empty")
	}
	if len(c.Notes) > 500 {
		return errors.New("notes cannot be longer than 500 characters")
	}

	return nil
}