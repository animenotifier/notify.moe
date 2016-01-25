'use strict'

let chalk = require('chalk')

const maxPopularAnime = 10

let updatePopularAnime = function() {
	console.log(chalk.yellow('✖'), 'Updating popular anime...')

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
		console.log(chalk.green('✔'), 'Updated popular anime.')
		return arn.set('Cache', 'popularAnime', popularAnime)
	})
}

arn.repeatedly(5 * 60, () => {
	updatePopularAnime()
})