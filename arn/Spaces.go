package arn

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

		// Initiate a client using DigitalOcean Spaces.
		Spaces, err = minio.New(endpoint, &minio.Options{
			Secure: true,
			Creds:  credentials.NewStaticV4(APIKeys.S3.ID, APIKeys.S3.Secret, ""),
		})

		if err != nil {
			log.Fatal(err)
		}
	}()
}
