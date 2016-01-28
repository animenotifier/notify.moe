'use strict'

let chalk = require('chalk')
let Promise = require('bluebird')

const animePageCacheTime = 120 * 60 * 1000

let updateAllAnimePages = Promise.coroutine(function*() {
	console.log(chalk.yellow('✖'), 'Updating all anime pages...')

	let now = new Date()

	for(let anime of arn.animeList) {
		if(anime.pageGenerated && now.getTime() - (new Date(anime.pageGenerated)).getTime() < animePageCacheTime)
			continue

		yield Promise.delay(250)
		arn.updateAnimePage(anime)
	}
})

arn.repeatedly(3 * 60 * 60, updateAllAnimePages)