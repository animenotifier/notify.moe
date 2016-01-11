'use strict'

let arn = require('../../../lib')

exports.get = function(request, response) {
	let user = request.user
	let term = request.params[0] || ''

	term = term.replace('%20', ' ').trim().toLowerCase()

	if(!term) {
		response.render({
			user
		})
		return
	}

	let animeList = []
	let hasSpace = term.indexOf(' ') !== -1

	arn.scan('Anime', function(anime) {
		if(!anime.title)
			return

		let title = anime.title.romaji.toLowerCase()

		if(title === term || (hasSpace && title.indexOf(term) !== -1)) {
			animeList.push(anime)
			return
		}

		let tryTitle = title => {
			let words = title.split(' ')
			for(let i = 0; i < words.length; i++) {
				let word = words[i]
				if(word === term || (term.length >= 2 && word.startsWith(term))) {
					animeList.push(anime)
					return true
				}
			}

			return false
		}

		if(tryTitle(title))
			return

		/*if(tryTitle(title.replace('ō', 'o').replace('Ō', 'o')))
			return

		if(tryTitle(title.replace('ō', 'ou').replace('Ō', 'ou')))
			return*/
	}, function() {
		animeList.sort((a, b) => a.title.romaji.localeCompare(b.title.romaji))

		response.render({
			user,
			term,
			animeList
		})
	})
}