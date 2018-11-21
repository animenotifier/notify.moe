package group

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Feed shows the group front page.
func Feed(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	id := ctx.Get("id")
	group, err := arn.GetGroup(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Group not found", err)
	}

	var member *arn.GroupMember

	if user != nil {
		member = group.FindMember(user.ID)
	}

	ctx.Data = getOpenGraph(ctx, group)
	return ctx.HTML(components.GroupFeed(group, member, user))
}
