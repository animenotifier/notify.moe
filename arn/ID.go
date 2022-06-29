package arn

import (
	"errors"
	"strings"

	"github.com/teris-io/shortid"
)

// ID is used for object identification and is simply a string.
type ID = string

// GenerateID generates a unique ID for a given collection.
func GenerateID(collection string) ID {
	id, _ := shortid.Generate()

	// Retry until we find an unused ID
	retry := 0

	for {
		_, err := DB.Get(collection, id)

		if err != nil && strings.Contains(err.Error(), "not found") {
			return id
		}

		retry++

		if retry > 10 {
			panic(errors.New("Can't generate unique ID"))
		}

		id, _ = shortid.Generate()
	}
}
