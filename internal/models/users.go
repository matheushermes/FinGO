package models

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
	"time"

	"github.com/matheushermes/FinGO/internal/security"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	FirstName string    `json:"firstName,omitempty"`
	LastName  string    `json:"lastName,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedIn time.Time `json:"createdIn,omitempty"`
}

func (u *User) IsValid(step string) error {
	u.trimSpaces()

	if err := u.checkBlankFields(); err != nil {
		return err
	}

	if err := u.validateEmail(); err != nil {
		return err
	}

	if err := u.validatePassword(); err != nil {
		return err
	}

	if err := u.hashPasswordIfNeeded(step); err != nil {
		return err
	}

	return nil
}

func (u *User) checkBlankFields() error {
	switch {
	case u.Username == "":
		return errors.New("username cannot be empty")
	case u.Email == "":
		return errors.New("email cannot be empty")
	case u.FirstName == "":
		return errors.New("first name cannot be empty")
	case u.LastName == "":
		return errors.New("last name cannot be empty")
	case u.Password == "":
		return errors.New("password cannot be empty")
	}
	return nil
}

func (u *User) trimSpaces() {
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
}

func (u *User) validateEmail() error {
	if _, err := mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email format")
	}
	return nil
}

func (u *User) validatePassword() error {
	pass := u.Password

	if len(pass) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if match, err := regexp.MatchString(`\d`, pass); err != nil || !match {
		return errors.New("password must contain at least one number")
	}
	if match, err := regexp.MatchString(`[A-Z]`, pass); err != nil || !match {
		return errors.New("password must contain at least one uppercase letter")
	}
	if match, err := regexp.MatchString(`[a-z]`, pass); err != nil || !match {
		return errors.New("password must contain at least one lowercase letter")
	}
	if match, err := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, pass); err != nil || !match {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func (u *User) hashPasswordIfNeeded(step string) error {
	if step == "register" {
		hashed, err := security.EncryptPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashed)
	}
	return nil
}
