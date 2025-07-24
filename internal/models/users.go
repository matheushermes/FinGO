package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID			uint64		`json:"id,omitempty"`
	Username	string		`json:"username,omitempty"`
	Email		string		`json:"email,omitempty"`
	FirstName	string		`json:"firstName,omitempty"`
	LastName	string		`json:"lastName,omitempty"`
	Password	string		`json:"password,omitempty"`
	CreatedIn 	time.Time	`json:"createdIn,omitempty"`
}

func (u *User) IsValid() error {
	var err error

	if err = u.checkBlankFields(); err != nil {
		return err
	}

	u.removeSpacesAtEnds()

	return nil
}

func (u *User) checkBlankFields() error {
	switch {
	case u.Username == "":
		return errors.New("username não pode ser vazio")
	case u.Email == "":
		return errors.New("email não pode ser vazio")
	case u.FirstName == "":
		return errors.New("firstName não pode ser vazio")
	case u.LastName == "":
		return errors.New("lastName não pode ser vazio")
	case u.Password == "":
		return errors.New("password não pode ser vazio")
	}

	return nil
}

func (u *User) removeSpacesAtEnds() {
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
	u.FirstName = strings.TrimSpace(u.FirstName)
	u.LastName = strings.TrimSpace(u.LastName)
}