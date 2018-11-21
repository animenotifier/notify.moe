package group

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

func getOpenGraph(ctx *aero.Context, group *arn.Group) *arn.OpenGraph {
	return &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     group.Name,
			"og:url":       "https://" + ctx.App.Config.Domain + group.Link(),
			"og:site_name": "notify.moe",
		},
	}
}
