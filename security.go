package main

import (
	"os"

	"github.com/aerogo/aero"
)

func configureHTTPS(app *aero.Application) {
	fullCertPath := "security/fullchain.pem"
	fullKeyPath := "security/privkey.pem"

	if _, err := os.Stat(fullCertPath); os.IsNotExist(err) {
		defaultCertPath := "security/default/fullchain.pem"
		err := os.Link(defaultCertPath, fullCertPath)

		if err != nil {
			panic(err)
		}
	}

	if _, err := os.Stat(fullKeyPath); os.IsNotExist(err) {
		defaultKeyPath := "security/default/privkey.pem"
		err := os.Link(defaultKeyPath, fullKeyPath)

		if err != nil {
			panic(err)
		}
	}

	// HTTPS
	app.Security.Load(fullCertPath, fullKeyPath)
}
