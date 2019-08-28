package amv

import (
	"strings"

	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
)

func getOpenGraph(amv *arn.AMV) *arn.OpenGraph {
	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     amv.Title.ByUser(nil) + " (AMV)",
			"og:url":       "https://" + assets.Domain + amv.Link(),
			"og:site_name": assets.Domain,
			"og:type":      "video.other",
		},
		Meta: map[string]string{},
	}

	openGraph.Tags["og:description"] = strings.Join(amv.Tags, ", ")

	if amv.File != "" {
		openGraph.Tags["og:video"] = amv.VideoLink()
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
