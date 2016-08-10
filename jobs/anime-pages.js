'use strict'

const animePageCacheTime = 120 * 60 * 1000

let updateAllAnimePages = coroutine(function*() {
	console.log(chalk.yellow('✖'), 'Updating all anime pages...')

	let now = new Date()

	for(let anime of arn.animeList) {
		if(anime.pageGenerated && now.getTime() - (new Date(anime.pageGenerated)).getTime() < animePageCacheTime)
			continue

		yield Promise.delay(1500)
		arn.updateAnimePage(anime)
	}
})

arn.repeatedly(3 * hours, updateAllAnimePages)