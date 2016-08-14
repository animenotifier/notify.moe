const maxPage = 5

let importAnimeFromAniList = coroutine(function*() {
	console.log(chalk.yellow('✖'), 'Import anime from anilist...')
	
	// Get an access token
	yield arn.listProviders.AniList.authorize()
	
	// Check the latest X pages for new anime edits
	for(let page = 1; page <= maxPage; page++) {
		yield Promise.delay(1200)
		
		// Get the list of 40 anime
		let animeList = yield arn.listProviders.AniList.getAnimeFromPage(page)
		
		// Write the new anime data into the DB
		yield animeList.map(anime => arn.set('Anime', anime.id, anime))

		console.log(chalk.green('✔'), `Finished importing anilist page ${chalk.yellow(page)} (${animeList.length} anime)`)
	}
	
	// Update the anime list used for the background jobs process
	arn.animeList = yield arn.all('Anime')
	
	console.log(chalk.green('✔'), `Finished importing basic anime data from AniList`)
})

arn.repeatedly(1 * hours, importAnimeFromAniList)