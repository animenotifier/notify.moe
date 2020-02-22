package filteranime

import (
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/validate"
)

// EndDate ...
func EndDate(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without a valid end date",
		func(anime *arn.Anime) bool {
			_, err := time.Parse(validate.DateFormat, anime.EndDate)
			return err != nil
		},
		nil,
	)
}
