const maxPopularAnime = 10

let updatePopularAnime = coroutine(function*() {
	console.log(chalk.yellow('✖'), 'Updating popular anime...')

	let popularAnime = arn.animeList.filter(anime => anime.watching)
	popularAnime.sort((a, b) => a.watching < b.watching ? 1 : -1)

	if(popularAnime.length > maxPopularAnime)
		popularAnime.length = maxPopularAnime

	yield arn.set('Cache', 'popularAnime', popularAnime)

	console.log(chalk.green('✔'), 'Updated popular anime.')
})

arn.repeatedly(5 * minutes, updatePopularAnime)