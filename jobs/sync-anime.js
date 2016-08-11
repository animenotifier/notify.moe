'use strict'

let importAnimeFromAniList = coroutine(function*() {
	console.log(chalk.yellow('✖'), 'Import anime from anilist...')

	yield arn.listProviders.AniList.authorize()

	let maxPage = 260
	for(let page = maxPage; page >= 1; page--) {
		yield Promise.delay(1200)

		let animeList = yield arn.listProviders.AniList.getAnimeFromPage(page)
		let tasks = animeList.map(anime => arn.set('Anime', anime.id, anime))
		yield Promise.all(tasks)

		console.log(chalk.green('✔'), `Finished importing anilist page ${chalk.yellow(page)} (${animeList.length} anime)`)
	}

	arn.animeList = yield arn.filter('Anime', anime => true)
})

arn.repeatedly(12 * hours, importAnimeFromAniList)