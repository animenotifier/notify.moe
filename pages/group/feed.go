package group

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
)

// Feed shows the group front page.
func Feed(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	id := ctx.Get("id")
	group, err := arn.GetGroup(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Group not found", err)
	}

	var member *arn.GroupMember

	if user != nil {
		member = group.FindMember(user.ID)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(group)
	return ctx.HTML(components.GroupFeed(group, member, user))
}
