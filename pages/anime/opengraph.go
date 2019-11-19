package anime

import (
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
)

func getOpenGraph(anime *arn.Anime) *arn.OpenGraph {
	description := anime.Summary

	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength-3] + "..."
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       anime.Title.Canonical,
			"og:image":       "https:" + anime.ImageLink("large"),
			"og:url":         "https://" + assets.Domain + anime.Link(),
			"og:site_name":   "notify.moe",
			"og:description": description,
		},
		Meta: map[string]string{
			"description": description,
			"keywords":    anime.Title.Canonical + ",anime",
		},
	}

	switch anime.Type {
	case "tv":
		openGraph.Tags["og:type"] = "video.tv_show"
	case "movie":
		openGraph.Tags["og:type"] = "video.movie"
	}

	return openGraph
}
