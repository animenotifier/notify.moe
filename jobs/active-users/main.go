package main

import (
	"fmt"
	"sort"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Caching list of active users")

	cache := arn.ListOfIDs{}

	// Filter out active users with an avatar
	users, err := arn.FilterUsers(func(user *arn.User) bool {
		return user.IsActive() && user.Avatar != ""
	})

	if err != nil {
		panic(err)
	}

	// Sort
	sort.Slice(users, func(i, j int) bool {
		return users[i].LastSeen > users[j].LastSeen
	})

	// Add users to list
	for _, user := range users {
		cache.IDList = append(cache.IDList, user.ID)
	}

	fmt.Println(len(cache.IDList), "users")

	err = arn.DB.Set("Cache", "active users", cache)

	if err != nil {
		panic(err)
	}

	color.Green("Finished.")
}
