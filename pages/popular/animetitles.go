package popular

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/stringutils"
)

// AnimeTitles returns a list of the 500 most popular anime titles.
func AnimeTitles(ctx aero.Context) error {
	maxLength, err := ctx.GetInt("count")

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Invalid value for count parameter", err)
	}

	popularAnimeTitles := []string{}
	popularAnime := arn.AllAnime()
	arn.SortAnimeByPopularity(popularAnime)

	if len(popularAnime) > maxLength {
		popularAnime = popularAnime[:maxLength]
	}

	for _, anime := range popularAnime {
		popularAnimeTitles = append(popularAnimeTitles, anime.Title.Canonical)

		if stringutils.ContainsUnicodeLetters(anime.Title.Japanese) {
			popularAnimeTitles = append(popularAnimeTitles, anime.Title.Japanese)
		}
	}

	return ctx.JSON(popularAnimeTitles)
}
