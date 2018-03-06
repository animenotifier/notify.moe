package group

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Edit ...
func Edit(ctx *aero.Context) string {
	id := ctx.Get("id")
	group, err := arn.GetGroup(id)
	user := utils.GetUser(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     group.Name,
			"og:url":       "https://" + ctx.App.Config.Domain + group.Link(),
			"og:site_name": "notify.moe",
		},
	}

	return ctx.HTML(components.GroupTabs(group, user) + editform.Render(group, "Edit group", user))
}
