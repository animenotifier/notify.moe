package popularanime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get search page.
func Get(ctx *aero.Context) string {
	// titleCount := 0
	// animeCount := 0

	// // let info: any = await bluebird.props({
	// // 	popular: arn.db.get('Cache', 'popularAnime'),
	// // 	stats: arn.db.get('Cache', 'animeStats')
	// // })

	// // return response.render({
	// // 	user,
	// // 	popularAnime: info.popular.anime,
	// // 	animeCount: info.stats.animeCount,
	// // 	titleCount: info.stats.titleCount,
	// // 	anime: null
	// // })

	// popular, _ := arn.GetPopularCache()

	// return ctx.HTML(components.Search(popular.Anime, titleCount, animeCount))
	animeList, err := arn.GetPopularAnimeCached()

	if err != nil {
		return ctx.HTML("There was a problem listing anime!")
	}

	return ctx.HTML(components.AnimeGrid(animeList))
}
