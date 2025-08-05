package repository

import (
	"database/sql"

	"github.com/matheushermes/FinGO/internal/models"
)

type cryptos struct {
	db *sql.DB
}

func NewCryptosRepository(db *sql.DB) *cryptos {
	return &cryptos{db}
}

func (c *cryptos) Create(crypto models.Crypto) (uint64, error) {
	query := `INSERT INTO cryptos (user_id, name, symbol, amount, purchase_price_usd, purchase_date, is_solid, notes)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`;
	var cryptoID uint64

	err := c.db.QueryRow(query, crypto.UserID, crypto.Name, crypto.Symbol, crypto.Amount, crypto.PurchasePriceUSD, crypto.PurchaseDate, crypto.IsSolid, crypto.Notes).Scan(&cryptoID)
	if err != nil {
		return 0, err
	}

	return cryptoID, nil
}