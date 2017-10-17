package groups

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	groups, err := arn.FilterGroups(func(group *arn.Group) bool {
		return !group.IsDraft
	})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching groups", err)
	}

	return ctx.HTML(components.Groups(groups, user))
}
