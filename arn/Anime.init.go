package arn

import "github.com/aerogo/api"

// AnimeSourceHumanReadable maps the anime source to a human readable version.
var AnimeSourceHumanReadable = map[string]string{}

// Register a list of supported anime status and source types.
func init() {
	DataLists["anime-types"] = []*Option{
		{"tv", "TV"},
		{"movie", "Movie"},
		{"ova", "OVA"},
		{"ona", "ONA"},
		{"special", "Special"},
		{"music", "Music"},
	}

	DataLists["anime-status"] = []*Option{
		{"current", "Current"},
		{"finished", "Finished"},
		{"upcoming", "Upcoming"},
		{"tba", "To be announced"},
	}

	DataLists["anime-sources"] = []*Option{
		{"", "Unknown"},
		{"original", "Original"},
		{"manga", "Manga"},
		{"novel", "Novel"},
		{"light novel", "Light novel"},
		{"visual novel", "Visual novel"},
		{"game", "Game"},
		{"book", "Book"},
		{"4-koma manga", "4-koma Manga"},
		{"music", "Music"},
		{"picture book", "Picture book"},
		{"web manga", "Web manga"},
		{"other", "Other"},
	}

	for _, option := range DataLists["anime-sources"] {
		AnimeSourceHumanReadable[option.Value] = option.Label
	}

	API.RegisterActions("Anime", []*api.Action{
		// Publish
		PublishAction(),

		// Unpublish
		UnpublishAction(),
	})
}
