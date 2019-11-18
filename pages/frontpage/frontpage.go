package frontpage

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
)

// Get ...
func Get(ctx aero.Context) error {
	description := "Anime list, tracker, database and notifier for new anime episodes. Create your own anime list and keep track of your progress as you watch."

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       assets.Manifest.Name,
			"og:description": description,
			"og:type":        "website",
			"og:url":         "https://" + assets.Domain,
			"og:image":       "https://" + assets.Domain + "/images/brand/220.png",
		},
		Meta: map[string]string{
			"description": description,
			"keywords":    "anime,list,tracker,notifier",
		},
	}

	return ctx.HTML(components.FrontPage())
}
