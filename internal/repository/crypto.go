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

func (c *cryptos) GetAllCryptos(userId uint64) ([]models.Crypto, error) {
	query := `SELECT * FROM cryptos WHERE user_id = $1`
	rows, err := c.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cryptos []models.Crypto
	for rows.Next() {
		var crypto models.Crypto
		if err := rows.Scan(&crypto.ID, &crypto.UserID, &crypto.Name, &crypto.Symbol, &crypto.Amount, &crypto.PurchasePriceUSD, &crypto.VariationPercent, &crypto.CurrentPriceUSD, &crypto.CurrentTotalValueUSD, &crypto.PurchaseDate, &crypto.IsSolid, &crypto.Notes, &crypto.CreatedAt, &crypto.UpdatedAt); err != nil {
			return nil, err
		}
		cryptos = append(cryptos, crypto)
	}

	return cryptos, nil
}

func (c *cryptos) GetCrypto(cryptoId uint64) (models.Crypto, error) {
	query := `SELECT * FROM cryptos WHERE id = $1`
	var crypto models.Crypto
	if err := c.db.QueryRow(query, cryptoId).Scan(&crypto.ID, &crypto.UserID, &crypto.Name, &crypto.Symbol, &crypto.Amount, &crypto.PurchasePriceUSD, &crypto.VariationPercent, &crypto.CurrentPriceUSD, &crypto.CurrentTotalValueUSD, &crypto.PurchaseDate, &crypto.IsSolid, &crypto.Notes, &crypto.CreatedAt, &crypto.UpdatedAt); err != nil {
		return models.Crypto{}, err
	}
	return crypto, nil
}

func (c *cryptos) UpdateCrypto(crypto models.Crypto) error {
	query := `UPDATE cryptos SET variation_percent = $1, current_price_usd = $2, current_total_value_usd =$3, updated_at = NOW() WHERE id = $4`
	_, err := c.db.Exec(query,
		crypto.VariationPercent,
		crypto.CurrentPriceUSD,
		crypto.CurrentTotalValueUSD,
		crypto.ID,
	)
	return err
}

func (c *cryptos) GetAllCryptosAllUsers() ([]models.Crypto, error) {
	query := `SELECT * FROM cryptos`
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cryptos []models.Crypto
	for rows.Next() {
		var crypto models.Crypto
		if err := rows.Scan(
			&crypto.ID, &crypto.UserID, &crypto.Name, &crypto.Symbol,
			&crypto.Amount, &crypto.PurchasePriceUSD, &crypto.VariationPercent,
			&crypto.CurrentPriceUSD, &crypto.CurrentTotalValueUSD,
			&crypto.PurchaseDate, &crypto.IsSolid, &crypto.Notes,
			&crypto.CreatedAt, &crypto.UpdatedAt,
		); err != nil {
			return nil, err
		}
		cryptos = append(cryptos, crypto)
	}
	return cryptos, nil
}