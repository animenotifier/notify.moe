package group

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	group, err := arn.GetGroup(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Group not found", err)
	}

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     group.Name,
			"og:url":       "https://" + ctx.App.Config.Domain + group.Link(),
			"og:site_name": "notify.moe",
		},
	}

	return ctx.HTML(components.Group(group))
}
