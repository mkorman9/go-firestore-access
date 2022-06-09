package main

import "github.com/mkorman9/go-commons/utils"

var clientCollection = "clients"

type ClientDocument struct {
	ID          string   `firestore:"id"`
	Gender      string   `firestore:"gender"`
	FirstName   string   `firestore:"firstName"`
	LastName    string   `firestore:"lastName"`
	Address     string   `firestore:"address"`
	PhoneNumber string   `firestore:"phoneNumber"`
	Email       string   `firestore:"email"`
	BirthDate   *int64   `firestore:"birthDate"`
	CreditCards []string `firestore:"creditCards"`
	IsDeleted   bool     `firestore:"isDeleted"`
}

func (d *ClientDocument) ToClient() *Client {
	return &Client{
		ID:          d.ID,
		Gender:      d.Gender,
		FirstName:   d.FirstName,
		LastName:    d.LastName,
		Address:     d.Address,
		PhoneNumber: d.PhoneNumber,
		Email:       d.Email,
		BirthDate:   utils.UnixPtrToTime(d.BirthDate),
		CreditCards: d.CreditCards,
		IsDeleted:   d.IsDeleted,
	}
}
