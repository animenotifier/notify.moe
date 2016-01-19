'use strict'

exports.get = function(request, response) {
	let genre = request.params[0]
	let genreSearch = `;${genre};`

	arn.filter('Anime', anime => anime.genres && (';' + anime.genres.join(';').toLowerCase() + ';').indexOf(genreSearch) !== -1).then(animeList => {
		animeList.sort((a, b) => {
			return a.startDate > b.startDate ? -1 : 1
		})

		response.render({
			genre,
			animeList
		})
	})
}