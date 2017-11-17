package frontpage

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	description := "Anime list, tracker, database and notifier for new anime episodes. Create your own anime list and keep track of your progress as you watch."

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       ctx.App.Config.Title,
			"og:description": description,
			"og:type":        "website",
			"og:url":         "https://" + ctx.App.Config.Domain,
			"og:image":       "https://" + ctx.App.Config.Domain + "/images/brand/220.png",
		},
		Meta: map[string]string{
			"description": description,
			"keywords":    "anime,list,tracker,notifier",
		},
	}

	return ctx.HTML(components.FrontPage())
}
