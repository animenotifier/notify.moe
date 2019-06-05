package group

import (
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
)

func getOpenGraph(group *arn.Group) *arn.OpenGraph {
	return &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       group.Name,
			"og:description": group.Tagline,
			"og:image":       "https:" + group.ImageLink("large"),
			"og:url":         "https://" + assets.Domain + group.Link(),
			"og:site_name":   "notify.moe",
		},
	}
}
