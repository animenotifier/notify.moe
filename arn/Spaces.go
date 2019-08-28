package arn

import (
	"log"

	"github.com/minio/minio-go/v6"
)

// Spaces represents our file storage server.
var Spaces *minio.Client

// initSpaces starts our file storage client.
func initSpaces() {
	if APIKeys.S3.ID == "" || APIKeys.S3.Secret == "" {
		return
	}

	go func() {
		var err error
		endpoint := "sfo2.digitaloceanspaces.com"
		ssl := true

		// Initiate a client using DigitalOcean Spaces.
		Spaces, err = minio.New(endpoint, APIKeys.S3.ID, APIKeys.S3.Secret, ssl)

		if err != nil {
			log.Fatal(err)
		}
	}()
}
