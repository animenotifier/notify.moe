package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// DuplicateMappings ...
func DuplicateMappings(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime with duplicate mappings",
		func(anime *arn.Anime) bool {
			existing := map[string]bool{}

			for _, mapping := range anime.Mappings {
				_, exists := existing[mapping.Service]

				if exists {
					return true
				}

				existing[mapping.Service] = true
			}

			return false
		},
		nil,
	)
}
