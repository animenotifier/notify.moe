package explorerelations

import (
	"net/http"
	"sort"

	"github.com/animenotifier/arn"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Sequels ...
func Sequels(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	animeList := user.AnimeList()
	sequels := []*utils.AnimeWithRelatedAnime{}

	for anime := range arn.StreamAnime() {
		item := animeList.Find(anime.ID)

		// Ignore if user added the anime and it's not "Planned" status
		if item != nil && item.Status != arn.AnimeListStatusPlanned {
			continue
		}

		prequels := anime.Prequels()

		for _, prequel := range prequels {
			item := animeList.Find(prequel.ID)

			if item != nil && item.Status == arn.AnimeListStatusCompleted {
				sequels = append(sequels, &utils.AnimeWithRelatedAnime{
					Anime:   anime,
					Related: prequel,
				})
				break
			}
		}
	}

	sort.Slice(sequels, func(i, j int) bool {
		aScore := sequels[i].Anime.Score()
		bScore := sequels[j].Anime.Score()

		if aScore == bScore {
			return sequels[i].Anime.Title.Canonical < sequels[j].Anime.Title.Canonical
		}

		return aScore > bScore
	})

	return ctx.HTML(components.ExploreAnimeSequels(sequels, user))
}
