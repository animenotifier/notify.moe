package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// DuplicateMappings ...
func DuplicateMappings(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Anime with duplicate mappings",
		func(anime *arn.Anime) bool {
			all := map[string]bool{}

			for _, mapping := range anime.Mappings {
				_, exists := all[mapping.Service]

				if exists {
					return true
				}

				all[mapping.Service] = true
			}

			return false
		},
		nil,
	)
}
