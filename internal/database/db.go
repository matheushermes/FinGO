package database

import (
	"database/sql"

	"github.com/matheushermes/FinGO/configs"
	_ "github.com/lib/pq" //PostgreSQL driver
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", configs.STRING_CONNECTION)
	if err != nil {
		return nil, err
	}
	
	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}