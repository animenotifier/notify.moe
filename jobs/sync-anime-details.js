let processQueue = coroutine(function*(queue) {
	for(let animeId of queue) {
		yield Promise.delay(1100)

		let details = yield arn.listProviders.AniList.getAnimeDetails(animeId)
		
		if(!details)
			continue
		
		yield arn.set('Anime', animeId, details)
		
		console.log(chalk.green('✔'), `Finished importing anime details of ${chalk.cyan(animeId)}`)
	}
})

let importAnimeDetailsFromAniList = coroutine(function*() {
	console.log(chalk.yellow('✖'), 'Import anime details from anilist...')
	
	// Get an access token
	yield arn.listProviders.AniList.authorize()

	// We have 2 queues: high and low priority
	let highPriority = []
	let lowPriority = []

	// Filter anime that have been imported already into the low priority queue
	arn.animeList.forEach(anime => {
		if(anime.description || anime.totalEpisodes || anime.duration)
			lowPriority.push(anime.id)
		else
			highPriority.push(anime.id)
	})

	console.log(chalk.yellow('✖'), highPriority.length, 'anime in high priority queue')
	console.log(chalk.yellow('✖'), lowPriority.length, 'anime in low priority queue')

	yield processQueue(highPriority)
	yield processQueue(lowPriority)
})

arn.repeatedly(12 * hours, importAnimeDetailsFromAniList)