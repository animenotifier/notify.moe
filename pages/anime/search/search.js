'use strict'

let arn = require('../../../lib')

exports.get = function(request, response) {
	let user = request.user
	let term = request.params[0]

	if(!term) {
		response.render({
			user
		})
		return
	}

	term = term.toLowerCase()

	let animeList = []
	arn.scan('Anime', function(anime) {
		if(!anime.title)
			return

		let tryTitle = title => {
			if(title.indexOf(term) !== -1 || term.indexOf(title) !== -1)
				animeList.push(anime)
		}

		let title = anime.title.toLowerCase()

		tryTitle(title)
		tryTitle(title.replace('ō', 'o').replace('Ō', 'o'))
		tryTitle(title.replace('ō', 'ou').replace('Ō', 'ou'))
	}, function() {
		animeList.sort((a, b) => a.title > b.title ? 1 : -1)

		response.render({
			user,
			term,
			animeList
		})
	})
}