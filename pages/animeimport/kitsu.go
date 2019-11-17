package animeimport

import (
	"fmt"
	"net/http"

	"github.com/akyoto/color"
	"github.com/animenotifier/kitsu"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/aerogo/aero"
)

// Kitsu anime import.
func Kitsu(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	kitsuAnimeObj, err := arn.Kitsu.Get("Anime", id)

	if kitsuAnimeObj == nil {
		return ctx.Error(http.StatusNotFound, "Kitsu anime not found", err)
	}

	kitsuAnime := kitsuAnimeObj.(*kitsu.Anime)

	// Convert
	anime, characters, relations := arn.NewAnimeFromKitsuAnime(kitsuAnime)

	// Add user ID to the anime
	anime.CreatedBy = user.ID

	// Save in database
	anime.Save()
	characters.Save()
	relations.Save()

	// Log
	fmt.Println(color.GreenString("✔"), anime.ID, anime.Title.Canonical)

	return nil
}
