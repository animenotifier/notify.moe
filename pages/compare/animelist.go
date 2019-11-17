package compare

import (
	"net/http"
	"sort"

	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// AnimeList ...
func AnimeList(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	nickA := ctx.Get("nick-1")
	nickB := ctx.Get("nick-2")

	a, err := arn.GetUserByNick(nickA)

	if err != nil || a == nil {
		return ctx.Error(http.StatusNotFound, "User not found: "+nickA, err)
	}

	b, err := arn.GetUserByNick(nickB)

	if err != nil || b == nil {
		return ctx.Error(http.StatusNotFound, "User not found: "+nickB, err)
	}

	comparisons := []*utils.Comparison{}
	countA := 0
	countB := 0

	animeListA := a.AnimeList()
	animeListB := b.AnimeList()

	for _, item := range animeListA.Items {
		if item.Status == arn.AnimeListStatusPlanned {
			continue
		}

		countA++

		comparisons = append(comparisons, &utils.Comparison{
			Anime: item.Anime(),
			ItemA: item,
			ItemB: animeListB.Find(item.AnimeID),
		})
	}

	for _, item := range animeListB.Items {
		if item.Status == arn.AnimeListStatusPlanned {
			continue
		}

		countB++

		if Contains(comparisons, item.AnimeID) {
			continue
		}

		comparisons = append(comparisons, &utils.Comparison{
			Anime: item.Anime(),
			ItemA: animeListA.Find(item.AnimeID),
			ItemB: item,
		})
	}

	sort.Slice(comparisons, func(i, j int) bool {
		aPopularity := comparisons[i].Anime.Popularity.Total()
		bPopularity := comparisons[j].Anime.Popularity.Total()

		if aPopularity == bPopularity {
			return comparisons[i].Anime.Title.Canonical < comparisons[j].Anime.Title.Canonical
		}

		return aPopularity > bPopularity
	})

	return ctx.HTML(components.CompareAnimeList(a, b, countA, countB, comparisons, user))
}

// Contains ...
func Contains(comparisons []*utils.Comparison, animeID string) bool {
	for _, comparison := range comparisons {
		if comparison.Anime.ID == animeID {
			return true
		}
	}

	return false
}
