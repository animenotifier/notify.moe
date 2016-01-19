'use strict'

let Promise = require('bluebird')
let repeatedly = require('../../../lib/utils/repeatedly')
let fs = Promise.promisifyAll(require('fs'))

let genres = {}

repeatedly(30 * 60, () => {
	console.log('Updating genre cache...')

	fs.readFileAsync('pages/anime/genres/genres.txt', 'utf8').then(genreText => {
		let genreList = genreText.split('\n')
		let tasks = []

		genreList.forEach(genre => {
			console.log(genre)

			genre = arn.fixGenre(genre)
			let genreSearch = `;${genre};`

			tasks.push(Promise.delay(tasks.length * 500).then(() => {
				console.log('Updating genre:', genre)

				let animeList = []

				return arn.forEach('Anime', anime => {
					if(!anime.genres)
						return

					if((';' + anime.genres.map(arn.fixGenre).join(';') + ';').indexOf(genreSearch) === -1)
						return

					animeList.push({
						id: anime.id,
						title: anime.title,
						startDate: anime.startDate,
						watching: anime.watching
					})
				}).then(() => {
					animeList.sort((a, b) => {
						if(!a.startDate)
							return 1

						if(!b.startDate)
							return -1

						return a.startDate > b.startDate ? -1 : 1
					})

					genres[genre] = animeList
				})
			}))
		})
	})
})

exports.get = (request, response) => {
	let genre = request.params[0]
	let animeList = genres[genre]

	response.render({
		genre,
		animeList
	})
}