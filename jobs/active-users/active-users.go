package main

import (
	"fmt"
	"sort"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Caching list of active users")

	// Filter out active users with an avatar
	users, err := arn.FilterUsers(func(user *arn.User) bool {
		return user.IsActive() && user.Avatar.Extension != ""
	})
	fmt.Println(len(users))

	arn.PanicOnError(err)

	// Sort
	sort.Slice(users, func(i, j int) bool {
		if users[i].LastSeen < users[j].LastSeen {
			return false
		}

		if users[i].LastSeen > users[j].LastSeen {
			return true
		}

		return users[i].Registered > users[j].Registered
	})

	// Add users to list
	SaveInCache("active users", users)

	// Sort by osu rank
	osuUsers := users[:]

	sort.Slice(osuUsers, func(i, j int) bool {
		return osuUsers[i].Accounts.Osu.PP > osuUsers[j].Accounts.Osu.PP
	})

	// Cut off users with 0 pp
	for index, user := range osuUsers {
		if user.Accounts.Osu.PP == 0 {
			osuUsers = osuUsers[:index]
			break
		}
	}

	// Save osu users
	SaveInCache("active osu users", osuUsers)

	// Sort by role
	staff := users[:]

	sort.Slice(staff, func(i, j int) bool {
		if staff[i].Role == "" {
			return false
		}

		if staff[j].Role == "" {
			return true
		}

		return staff[i].Role == "admin"
	})

	// Cut off non-staff
	for index, user := range staff {
		if user.Role == "" {
			staff = staff[:index]
			break
		}
	}

	// Save staff users
	SaveInCache("active staff users", staff)

	// Sort by anime watching list length
	watching := users[:]

	sort.Slice(watching, func(i, j int) bool {
		return len(watching[i].AnimeList().FilterStatus(arn.AnimeListStatusWatching).Items) > len(watching[j].AnimeList().FilterStatus(arn.AnimeListStatusWatching).Items)
	})

	// Save watching users
	SaveInCache("active anime watching users", watching)

	color.Green("Finished.")
}

// SaveInCache ...
func SaveInCache(key string, users []*arn.User) {
	cache := arn.ListOfIDs{
		IDList: GenerateIDList(users),
	}

	fmt.Println(len(cache.IDList), key)
	arn.PanicOnError(arn.DB.Set("Cache", key, cache))
}

// GenerateIDList generates an ID list from a slice of users.
func GenerateIDList(users []*arn.User) []string {
	list := []string{}

	for _, user := range users {
		list = append(list, user.ID)
	}

	return list
}
