package users

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Active ...
func Active(ctx *aero.Context) string {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.IsActive() && user.HasAvatar()
	})

	arn.SortUsersLastSeen(users)

	return ctx.HTML(components.Users(users))
}

// Osu ...
func Osu(ctx *aero.Context) string {
	users, err := arn.GetListOfUsersCached("active osu users")

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not fetch user data", err)
	}

	if len(users) > 50 {
		users = users[:50]
	}

	return ctx.HTML(components.OsuRankingList(users))
}

// Staff ...
func Staff(ctx *aero.Context) string {
	users, err := arn.GetListOfUsersCached("active staff users")

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not fetch user data", err)
	}

	return ctx.HTML(components.Users(users))
}

// AnimeWatching ...
func AnimeWatching(ctx *aero.Context) string {
	users, err := arn.GetListOfUsersCached("active anime watching users")

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not fetch user data", err)
	}

	return ctx.HTML(components.Users(users))
}
