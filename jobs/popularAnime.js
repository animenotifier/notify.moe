'use strict'

const maxPopularAnime = 10

let updatePopularAnime = function() {
	console.log('Updating popular anime...')

	let popularAnime = []

	arn.forEach('Anime', anime => {
		if(anime.watching)
			popularAnime.push(anime)
	}).then(() => {
		popularAnime.sort((a, b) => a.watching < b.watching ? 1 : -1)

		if(popularAnime.length > maxPopularAnime)
			popularAnime.length = maxPopularAnime

		return popularAnime
	}).then(popularAnime => {
		console.log('Updated popular anime.')
		return arn.set('Cache', 'popularAnime', popularAnime)
	})
}

arn.repeatedly(5 * 60, () => {
	arn.cacheLimiter.removeTokens(1, updatePopularAnime)
})