package group

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Edit ...
func Edit(ctx aero.Context) error {
	id := ctx.Get("id")
	group, err := arn.GetGroup(id)
	user := arn.GetUserFromContext(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Group not found", err)
	}

	var member *arn.GroupMember

	if user != nil {
		member = group.FindMember(user.ID)
	}

	return ctx.HTML(components.GroupHeader(group, member, user) + editform.Render(group, "Edit group", user))
}

// EditImage renders the form to edit the group images.
func EditImage(ctx aero.Context) error {
	id := ctx.Get("id")
	group, err := arn.GetGroup(id)
	user := arn.GetUserFromContext(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Group not found", err)
	}

	var member *arn.GroupMember

	if user != nil {
		member = group.FindMember(user.ID)
	}

	return ctx.HTML(components.EditGroupImage(group, member, user))
}
