package popularanime

import (
	"github.com/aerogo/aero"
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
	return ctx.HTML("Coming soon.")
}
