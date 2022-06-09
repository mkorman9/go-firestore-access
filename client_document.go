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

func (e *ClientDocument) ToClient() *Client {
	return &Client{
		ID:          e.ID,
		Gender:      e.Gender,
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		Address:     e.Address,
		PhoneNumber: e.PhoneNumber,
		Email:       e.Email,
		BirthDate:   utils.UnixPtrToTime(e.BirthDate),
		CreditCards: e.CreditCards,
		IsDeleted:   e.IsDeleted,
	}
}
