package amv

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

func getOpenGraph(ctx *aero.Context, amv *arn.AMV) *arn.OpenGraph {
	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     amv.Title.ByUser(nil) + " (AMV)",
			"og:url":       "https://" + ctx.App.Config.Domain + amv.Link(),
			"og:site_name": ctx.App.Config.Domain,
			"og:type":      "video.other",
		},
		Meta: map[string]string{},
	}

	// if amv.MainAnime() != nil {
	// 	openGraph.Tags["og:image"] = amv.MainAnime().ImageLink("large")
	// 	openGraph.Tags["og:description"] = amv.MainAnime().Title.Canonical + " (" + strings.Join(amv.Tags, ", ") + ")"
	// } else {
	// 	openGraph.Tags["og:description"] = strings.Join(amv.Tags, ", ")
	// }

	openGraph.Tags["og:description"] = strings.Join(amv.Tags, ", ")

	if amv.File != "" {
		openGraph.Tags["og:video"] = "https://" + ctx.App.Config.Domain + "/videos/amvs/" + amv.File
		openGraph.Tags["og:video:type"] = "video/webm"
		openGraph.Tags["og:video:width"] = "640"
		openGraph.Tags["og:video:height"] = "360"

		openGraph.Meta["twitter:player"] = openGraph.Tags["og:video"]
		openGraph.Meta["twitter:player:width"] = openGraph.Tags["og:video:width"]
		openGraph.Meta["twitter:player:height"] = openGraph.Tags["og:video:height"]
		openGraph.Meta["twitter:player:stream"] = openGraph.Tags["og:video"]
		openGraph.Meta["twitter:player:stream:content_type"] = openGraph.Tags["og:video:type"]
	}

	return openGraph
}
