package filteranime

import (
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/validate"
)

// StartDate ...
func StartDate(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without a valid start date",
		func(anime *arn.Anime) bool {
			_, err := time.Parse(validate.DateFormat, anime.StartDate)
			return err != nil
		},
		nil,
	)
}
