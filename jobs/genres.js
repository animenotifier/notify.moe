let updateGenres = coroutine(function*() {
	console.log(chalk.cyan('↻'), 'Updating genre cache...')

	let genreText = yield fs.readFileAsync('pages/anime/genres/genres.txt', 'utf8')
	let genreList = genreText.split('\n')

	for(let genre of genreList) {
		genre = arn.fixGenre(genre)
		let genreSearch = `;${genre};`

		yield Promise.delay(1000)

		let animeList = arn.animeList.filter(anime => {
			if(!anime.watching)
				return false

			if(!anime.genres)
				return false

			if((';' + anime.genres.map(arn.fixGenre).join(';') + ';').indexOf(genreSearch) === -1)
				return false

			return true
		})
		
		// Reduce data size
		animeList = animeList.map(anime => {
			return {
				id: anime.id,
				title: anime.title,
				image: anime.image ? anime.image : '',
				watching: anime.watching,
				startDate: anime.startDate
			}
		})

		animeList.sort((a, b) => {
			if(a.watching === b.watching) {
				if(!a.startDate)
					return 1

				if(!b.startDate)
					return -1

				return a.startDate > b.startDate ? -1 : 1
			} else if(!a.watching) {
				return 1
			} else if(!b.watching) {
				return -1
			} else {
				return a.watching > b.watching ? -1 : 1
			}
		})

		yield arn.db.set('Genres', genre, {
			genre,
			animeList
		})

		console.log(chalk.green('✔'), `Updated genre ${chalk.yellow(genre)} (${animeList.length} anime)`)
	}
})

arn.repeatedly(30 * minutes, updateGenres)
