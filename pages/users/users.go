package users

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Active ...
func Active(ctx *aero.Context) string {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.IsActive() && user.HasAvatar()
	})

	sort.Slice(users, func(i, j int) bool {
		followersA := users[i].FollowersCount()
		followersB := users[j].FollowersCount()

		if followersA == followersB {
			return users[i].Nick < users[j].Nick
		}

		return followersA > followersB
	})

	// arn.SortUsersLastSeen(users)

	return ctx.HTML(components.Users(users))
}

// Osu ...
func Osu(ctx *aero.Context) string {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.IsActive() && user.HasAvatar() && user.Accounts.Osu.PP > 0
	})

	// Sort by pp
	sort.Slice(users, func(i, j int) bool {
		return users[i].Accounts.Osu.PP > users[j].Accounts.Osu.PP
	})

	if len(users) > 50 {
		users = users[:50]
	}

	return ctx.HTML(components.OsuRankingList(users))
}

// Staff ...
func Staff(ctx *aero.Context) string {
	users := arn.FilterUsers(func(user *arn.User) bool {
		return user.IsActive() && user.HasAvatar() && user.Role != ""
	})

	sort.Slice(users, func(i, j int) bool {
		if users[i].Role == "" {
			return false
		}

		if users[j].Role == "" {
			return true
		}

		return users[i].Role == "admin"
	})

	return ctx.HTML(components.Users(users))
}
