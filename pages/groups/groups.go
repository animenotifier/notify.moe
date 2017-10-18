package groups

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const groupsPerPage = 12

// Get ...
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	groups, err := arn.FilterGroups(func(group *arn.Group) bool {
		return !group.IsDraft
	})

	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i].Members) == len(groups[j].Members) {
			return groups[i].Created > groups[j].Created
		}

		return len(groups[i].Members) > len(groups[j].Members)
	})

	if len(groups) > groupsPerPage {
		groups = groups[:groupsPerPage]
	}

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching groups", err)
	}

	return ctx.HTML(components.Groups(groups, groupsPerPage, user))
}
