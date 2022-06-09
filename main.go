package main

import (
	"context"
	"encoding/json"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"github.com/mkorman9/go-commons/firestorelib"
	"github.com/mkorman9/go-commons/logging"
	"github.com/mkorman9/go-commons/utils"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"time"
)

func main() {
	config.AddDriver(yaml.Driver)
	_ = config.LoadFiles("config.yml")

	logging.Setup()

	fs, closeFS, err := firestorelib.NewClient()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to Firestore")
	}
	defer closeFS()

	id := uuid.NewV4().String()
	client := Client{
		ID:          id,
		Gender:      GenderMale,
		FirstName:   "AAA",
		LastName:    "BBB",
		Address:     "AAA 123/456",
		PhoneNumber: "123-456-789",
		Email:       "aaa@example.com",
		BirthDate:   utils.TimePtr(time.Now().UTC()),
		CreditCards: []string{"1111 2222 3333 4444"},
		IsDeleted:   false,
	}
	if _, err := fs.Collection(clientCollection).Doc(id).Set(context.Background(), client.ToDocument()); err != nil {
		log.Fatal().Err(err).Msg("Failed to save document")
	}

	documentSnapshot, err := fs.Collection(clientCollection).Doc(id).Get(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to select record")
	}

	var document ClientDocument
	err = documentSnapshot.DataTo(&document)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to convert document to struct")
	}

	b, _ := json.Marshal(document.ToClient())
	log.Info().Msg(string(b))
}
