package export

import (
	"bytes"
	"fmt"

	"github.com/aerogo/aero"
)

// CSV renders the anime list items in plain text format.
func CSV(ctx aero.Context) error {
	animeList, err := getAnimeList(ctx)

	if err != nil {
		return err
	}

	buffer := bytes.Buffer{}

	// Header
	buffer.WriteString("Title,Status,Episodes,Overall,Story,Visuals,Soundtrack,Rewatched\n")

	// List items
	for _, item := range animeList.Items {
		anime := item.Anime()
		fmt.Fprintf(
			&buffer,
			"%s,%s,%d,%.1f,%.1f,%.1f,%.1f,%d\n",
			anime.Title.Canonical,
			item.Status,
			item.Episodes,
			item.Rating.Overall,
			item.Rating.Story,
			item.Rating.Visuals,
			item.Rating.Soundtrack,
			item.RewatchCount,
		)
	}

	ctx.Response().SetHeader("Content-Type", "text/csv")
	ctx.Response().SetHeader("Content-Disposition", "attachment; filename=\"anime-list.csv\"")
	return ctx.String(buffer.String())
}
