package compare

import (
	"net/http"
	"sort"

	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// AnimeList ...
func AnimeList(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
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

	for _, item := range a.AnimeList().Items {
		if item.Status == arn.AnimeListStatusPlanned {
			continue
		}

		comparisons = append(comparisons, &utils.Comparison{
			Anime: item.Anime(),
			ItemA: item,
			ItemB: b.AnimeList().Find(item.AnimeID),
		})
	}

	for _, item := range b.AnimeList().Items {
		if Contains(comparisons, item.AnimeID) || item.Status == arn.AnimeListStatusPlanned {
			continue
		}

		comparisons = append(comparisons, &utils.Comparison{
			Anime: item.Anime(),
			ItemA: a.AnimeList().Find(item.AnimeID),
			ItemB: item,
		})
	}

	sort.Slice(comparisons, func(i, j int) bool {
		return comparisons[i].Anime.Popularity.Total() > comparisons[j].Anime.Popularity.Total()
	})

	return ctx.HTML(components.CompareAnimeList(a, b, comparisons, user))
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
