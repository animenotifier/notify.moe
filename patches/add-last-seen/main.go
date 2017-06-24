package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	// Get a stream of all users
	allUsers, err := arn.AllUsers()

	if err != nil {
		panic(err)
	}

	// Iterate over the stream
	for user := range allUsers {
		if user.LastSeen != "" {
			continue
		}

		user.LastSeen = user.LastLogin

		if user.LastSeen == "" {
			user.LastSeen = user.Registered
		}

		user.Save()
	}
}
