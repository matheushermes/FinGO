package models

import "time"

type PriceAlert struct {
    ID                  uint64    `json:"id,omitempty"`
    UserEmail           string    `json:"user_email,omitempty"`
    CryptoID            uint64    `json:"crypto_id,omitempty"`
    Symbol              string    `json:"symbol,omitempty"`
    PercentageChange    float64   `json:"percentage_change,omitempty"`
    Direction           string    `json:"direction,omitempty"`
    TargetPrice         float64   `json:"target_price,omitempty"`
    Triggered           bool      `json:"triggered,omitempty"`
    CreatedAt           time.Time `json:"created_at,omitempty"`
}
