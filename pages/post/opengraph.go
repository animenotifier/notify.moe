package post

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
)

func getOpenGraph(ctx *aero.Context, post *arn.Post) *arn.OpenGraph {
	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       post.TitleByUser(nil),
			"og:description": utils.CutLongDescription(post.Text),
			"og:url":         "https://" + ctx.App.Config.Domain + post.Link(),
			"og:site_name":   ctx.App.Config.Domain,
			"og:type":        "article",
		},
	}

	return openGraph
}
