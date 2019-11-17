package groups

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Joined shows the most popular joined groups.
func Joined(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	groups := fetchGroups(user.ID)

	// Sort by join date
	sort.Slice(groups, func(i, j int) bool {
		aMember := groups[i].FindMember(user.ID)
		bMember := groups[j].FindMember(user.ID)

		return aMember.Joined > bMember.Joined
	})

	return render(ctx, groups)
}
