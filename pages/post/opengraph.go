package post

import (
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/utils"
)

func getOpenGraph(post *arn.Post) *arn.OpenGraph {
	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       post.TitleByUser(nil),
			"og:description": utils.CutLongDescription(post.Text),
			"og:url":         "https://" + assets.Domain + post.Link(),
			"og:site_name":   assets.Domain,
			"og:type":        "article",
		},
	}

	return openGraph
}
