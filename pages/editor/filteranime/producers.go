package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Producers ...
func Producers(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without producers",
		func(anime *arn.Anime) bool {
			return len(anime.ProducerIDs) == 0
		},
		nil,
	)
}
