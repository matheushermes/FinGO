package models

import "time"

type User struct {
	ID			uint64		`json:"id,omitempty"`
	Username	string		`json:"username,omitempty"`
	Email		string		`json:"email,omitempty"`
	FirstName	string		`json:"firstName,omitempty"`
	LastName	string		`json:"lastName,omitempty"`
	Password	string		`json:"password,omitempty"`
	CreatedIn 	time.Time	`json:"createdIn,omitempty"`
}