package models

import "time"

type Crypto struct {
	ID uint64 					`json:"id,omitempty"`
	UserID uint64 				`json:"userID,omitempty"`
	Name string 				`json:"name,omitempty"`
	Symbol string 				`json:"symbol,omitempty"`
	Amount float64 				`json:"amount,omitempty"`
	PurchasePriceUSD float64 	`json:"purchasePriceUSD,omitempty"`
	PurchaseDate time.Time 		`json:"purchaseDate,omitempty"`
	IsSolid bool 				`json:"isSolid,omitempty"`
	Notes string 				`json:"notes,omitempty"`
	CreatedAt time.Time 		`json:"createdAt,omitempty"`
	UpdatedAt time.Time 		`json:"updatedAt,omitempty"`
}

func (c *Crypto) ValidationsCryptos() error {
	return nil
}