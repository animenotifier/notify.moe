package halloffame

import (
	"sort"
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const minYear = 1963

// Get ...
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	maxYear := time.Now().Year() - 1
	hallOfFameEntries := []*utils.HallOfFameEntry{}

	animes := arn.FilterAnime(func(anime *arn.Anime) bool {
		if len(anime.StartDate) < 4 {
			return false
		}

		year := anime.StartDateTime().Year()

		if year > maxYear || year < minYear {
			return false
		}

		if anime.Status != "finished" {
			return false
		}

		if anime.Type != "tv" {
			return false
		}

		return true
	})

	arn.SortAnimeByQuality(animes)

	yearsAdded := map[int]bool{}

	for _, anime := range animes {
		year := anime.StartDateTime().Year()

		_, exists := yearsAdded[year]

		if exists {
			continue
		}

		hallOfFameEntries = append(hallOfFameEntries, &utils.HallOfFameEntry{
			Year:  year,
			Anime: anime,
		})

		yearsAdded[year] = true
	}

	sort.Slice(hallOfFameEntries, func(i, j int) bool {
		return hallOfFameEntries[i].Year > hallOfFameEntries[j].Year
	})

	return ctx.HTML(components.HallOfFame(hallOfFameEntries, user))
}
