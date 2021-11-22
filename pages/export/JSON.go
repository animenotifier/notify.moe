package export

import (
	"github.com/aerogo/aero"
)

// JSON renders the anime list items in JSON format.
func JSON(ctx aero.Context) error {
	animeList, err := getAnimeList(ctx)

	if err != nil {
		return err
	}

	ctx.Response().SetHeader("Content-Disposition", "attachment; filename=\"anime-list.json\"")
	return ctx.JSON(animeList)
}
