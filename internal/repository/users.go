package repository

import (
	"database/sql"

	"github.com/matheushermes/FinGO/internal/models"
)

type users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *users {
	return &users{db}
}

func (u *users) Create(user models.User) (uint64, error) {
	query := `INSERT INTO users (username, email, first_name, last_name, password) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var userID uint64

	err := u.db.QueryRow(query, user.Username, user.Email, user.FirstName, user.LastName, user.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (u *users) FindByEmail(email string) (models.User, error) {
	query := `SELECT id, password FROM users WHERE email = $1`
	var user models.User
	
	err := u.db.QueryRow(query, email).Scan(&user.ID, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}