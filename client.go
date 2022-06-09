package main

import (
	"github.com/mkorman9/go-commons/utils"
	"time"
)

type Gender = string

const (
	GenderUndefined = Gender("-")
	GenderMale      = Gender("M")
	GenderFemale    = Gender("F")
)

type Client struct {
	ID          string     `json:"id"`
	Gender      Gender     `json:"gender"`
	FirstName   string     `json:"firstName"`
	LastName    string     `json:"lastName"`
	Address     string     `json:"address,omitempty"`
	PhoneNumber string     `json:"phoneNumber,omitempty"`
	Email       string     `json:"email,omitempty"`
	BirthDate   *time.Time `json:"birthDate,omitempty"`
	CreditCards []string   `json:"creditCards"`
	IsDeleted   bool       `json:"-"`
}

func (c *Client) ToDocument() *ClientDocument {
	return &ClientDocument{
		ID:          c.ID,
		Gender:      c.Gender,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		Address:     c.Address,
		PhoneNumber: c.PhoneNumber,
		Email:       c.Email,
		BirthDate:   utils.TimePtrToUnix(c.BirthDate),
		CreditCards: c.CreditCards,
		IsDeleted:   c.IsDeleted,
	}
}
