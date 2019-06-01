package filteranime

import (
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// StartDate ...
func StartDate(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without a valid start date",
		func(anime *arn.Anime) bool {
			_, err := time.Parse(arn.AnimeDateFormat, anime.StartDate)
			return err != nil
		},
		nil,
	)
}
