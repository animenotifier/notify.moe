'use strict'

let fs = require('fs')
let AutoUpdateCache = require('../../../lib/classes/AutoUpdateCache')

let genres = new AutoUpdateCache(() => {
	console.log('Updating genre cache...')

	let allGenres = {}
	let genreList = fs.readFileSync('pages/anime/genres/genres.txt', 'utf8').split('\n')
	let tasks = []

	genreList.forEach(genre => {
		console.log('Updating cache for genre:', genre)

		genre = genre.toLowerCase()
		let genreSearch = `;${genre};`

		tasks.push(arn.filter('Anime', anime => anime.genres && (';' + anime.genres.join(';').toLowerCase() + ';').indexOf(genreSearch) !== -1).then(animeList => {
			animeList.sort((a, b) => {
				if(!a.startDate)
					return 1

				if(!b.startDate)
					return -1

				return a.startDate > b.startDate ? -1 : 1
			})

			allGenres[genre] = animeList
		}))
	})

	return Promise.all(tasks).then(() => allGenres)
}, 30 * 60, {})

exports.get = (request, response) => {
	let genre = request.params[0]
	let animeList = genres.cache[genre]

	response.render({
		genre,
		animeList
	})
}