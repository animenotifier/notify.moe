package thread

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
)

func getOpenGraph(ctx *aero.Context, thread *arn.Thread) *arn.OpenGraph {
	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       thread.Title,
			"og:description": utils.CutLongDescription(thread.Text),
			"og:url":         "https://" + ctx.App.Config.Domain + thread.Link(),
			"og:site_name":   ctx.App.Config.Domain,
			"og:type":        "article",
		},
	}

	return openGraph
}
