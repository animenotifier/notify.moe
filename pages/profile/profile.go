package profile

import (
	"sort"
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/utils"
)

const (
	maxCharacters = 6
	maxFriends    = 7
	maxStudios    = 4
)

// Get user profile page.
func Get(ctx aero.Context) error {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(404, "User not found", err)
	}

	return Profile(ctx, viewUser)
}

// Profile renders the user profile page of the given viewUser.
func Profile(ctx aero.Context, viewUser *arn.User) error {
	user := utils.GetUser(ctx)

	// Anime list
	animeList := viewUser.AnimeList()

	if user == nil || user.ID != viewUser.ID {
		animeList = animeList.WithoutPrivateItems()
	}

	completedList := animeList.FilterStatus(arn.AnimeListStatusCompleted)
	completedList.SortByRating()

	// Genres
	topGenres := animeList.TopGenres(5)

	// Studios
	animeWatchingTime := time.Duration(0)
	studios := map[string]float64{}
	var topStudios []*arn.Company

	for _, item := range animeList.Items {
		if item.Status != arn.AnimeListStatusCompleted {
			continue
		}

		currentWatch := item.Episodes * item.Anime().EpisodeLength
		reWatch := item.RewatchCount * item.Anime().EpisodeCount * item.Anime().EpisodeLength
		duration := time.Duration(currentWatch + reWatch)
		animeWatchingTime += duration * time.Minute
		rating := 0.0

		if item.Rating.Overall != 0 {
			rating = item.Rating.Overall - arn.AverageRating
		} else {
			// Add 0.1 to avoid all affinities being 0 when a user doesn't have any rated anime.
			rating = 0.1
		}

		for _, studio := range item.Anime().Studios() {
			affinity, exists := studios[studio.ID]

			if !exists {
				topStudios = append(topStudios, studio)
			}

			studios[studio.ID] = affinity + rating
		}
	}

	sort.Slice(topStudios, func(i, j int) bool {
		affinityA := studios[topStudios[i].ID]
		affinityB := studios[topStudios[j].ID]

		if affinityA == affinityB {
			return topStudios[i].Name.English < topStudios[j].Name.English
		}

		return affinityA > affinityB
	})

	if len(topStudios) > maxStudios {
		topStudios = topStudios[:maxStudios]
	}

	// Open graph
	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":         viewUser.Nick,
			"og:image":         viewUser.AvatarLink("large"),
			"og:url":           "https://" + assets.Domain + viewUser.Link(),
			"og:site_name":     "notify.moe",
			"og:description":   utils.CutLongDescription(viewUser.Introduction),
			"og:type":          "profile",
			"profile:username": viewUser.Nick,
		},
		Meta: map[string]string{
			"description": utils.CutLongDescription(viewUser.Introduction),
			"keywords":    viewUser.Nick + ",profile",
		},
	}

	// Friends
	friends := viewUser.Follows().UsersWhoFollowBack()

	arn.SortUsersFollowers(friends)

	if len(friends) > maxFriends {
		friends = friends[:maxFriends]
	}

	// Activities
	activities := arn.FilterActivities(func(activity arn.Activity) bool {
		return activity.GetCreatedBy() == viewUser.ID
	})

	// Time zone offset
	var timeZoneOffset time.Duration
	analytics := viewUser.Analytics()

	if analytics != nil {
		timeZoneOffset = time.Duration(-analytics.General.TimezoneOffset) * time.Minute
	}

	now := time.Now().UTC().Add(timeZoneOffset)
	weekDay := int(now.Weekday())
	currentYearDay := int(now.YearDay())

	// Day offset is the number of days we need to reach Sunday
	dayOffset := 0

	if weekDay > 0 {
		dayOffset = 7 - weekDay
	}

	dayToActivityCount := map[int]int{}

	for _, activity := range activities {
		activityTime := activity.GetCreatedTime().Add(timeZoneOffset)
		activityYearDay := activityTime.YearDay()
		days := currentYearDay - activityYearDay
		dayToActivityCount[days+dayOffset]++
	}

	// Characters
	characters := []*arn.Character{}

	for character := range arn.StreamCharacters() {
		if arn.Contains(character.Likes, viewUser.ID) {
			characters = append(characters, character)
		}
	}

	sort.Slice(characters, func(i, j int) bool {
		aLikes := len(characters[i].Likes)
		bLikes := len(characters[j].Likes)

		if aLikes == bLikes {
			return characters[i].Name.Canonical < characters[j].Name.Canonical
		}

		return aLikes > bLikes
	})

	if len(characters) > maxCharacters {
		characters = characters[:maxCharacters]
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = openGraph

	return ctx.HTML(components.Profile(
		viewUser,
		user,
		animeList,
		completedList,
		characters,
		friends,
		topGenres,
		topStudios,
		animeWatchingTime,
		dayToActivityCount,
		ctx.Path(),
	))
}
