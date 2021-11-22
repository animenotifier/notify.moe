package export

import (
	"bytes"
	"fmt"

	"github.com/aerogo/aero"
)

// TXT renders the anime list items in plain text format.
func TXT(ctx aero.Context) error {
	animeList, err := getAnimeList(ctx)

	if err != nil {
		return err
	}

	buffer := bytes.Buffer{}

	for _, item := range animeList.Items {
		anime := item.Anime()
		fmt.Fprintf(&buffer, "Title: %s\n", anime.Title.Canonical)
		fmt.Fprintf(&buffer, "Status: %s\n", item.Status)
		fmt.Fprintf(&buffer, "Episodes: %d\n", item.Episodes)
		fmt.Fprintf(&buffer, "Overall: %.1f\n", item.Rating.Overall)
		fmt.Fprintf(&buffer, "Story: %.1f\n", item.Rating.Story)
		fmt.Fprintf(&buffer, "Visuals: %.1f\n", item.Rating.Visuals)
		fmt.Fprintf(&buffer, "Soundtrack: %.1f\n", item.Rating.Soundtrack)
		fmt.Fprintf(&buffer, "Rewatched: %d\n", item.RewatchCount)
		fmt.Fprintf(&buffer, "Notes: %s\n", item.Notes)
		buffer.WriteString("\n")
	}

	ctx.Response().SetHeader("Content-Disposition", "attachment; filename=\"anime-list.txt\"")
	return ctx.Text(buffer.String())
}
