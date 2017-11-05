package group

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Forum ...
func Forum(ctx *aero.Context) string {
	id := ctx.Get("id")
	group, err := arn.GetGroup(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Group not found", err)
	}

	return ctx.HTML(components.GroupForum(group))
}
