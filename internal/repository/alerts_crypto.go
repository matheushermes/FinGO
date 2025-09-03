package repository

import (
	"database/sql"

	"github.com/matheushermes/FinGO/internal/models"
)

type alert struct {
	db *sql.DB
}

func NewAlertsRepository(db *sql.DB) *alert {
	return &alert{db}
}

func (a *alert) CreateAlert(alert models.PriceAlert) (uint64, error) {
	query := `
		INSERT INTO price_alerts 
		(user_email, crypto_id, symbol, percentage_change, direction, target_price) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	var alertID uint64
	err := a.db.QueryRow(
		query,
		alert.UserEmail,
		alert.CryptoID,
		alert.Symbol,
		alert.PercentageChange,
		alert.Direction,
		alert.TargetPrice,
	).Scan(&alertID)

	if err != nil {
		return 0, err
	}

	return alertID, nil
}
