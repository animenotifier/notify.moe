'use strict'

const animePageCacheTime = 60 * 60 * 1000

let updateAllAnimePages = () => {
	let now = new Date()

	return arn.forEach('Anime', anime => {
		if(anime.pageGenerated && now.getTime() - (new Date(anime.pageGenerated)).getTime() < animePageCacheTime)
			return

		arn.cacheLimiter.removeTokens(1, () => {
			arn.updateAnimePage(anime)
		})
	})
}

arn.repeatedly(3 * 60 * 60, updateAllAnimePages)