'use strict'

let arn = require('../../lib')

let animeToIdJSONString = null

exports.get = function(request, response) {
	let user = request.user
	let animeId = parseInt(request.params[0])

	if(animeId) {
		arn.get('Anime', animeId).then(anime => {
			response.render({
				user,
				anime
			})
		}).catch(error => {
			response.writeHead(404)
			response.end('Anime not found.')
		})
		return
	}

	if(animeToIdJSONString === null) {
		let animeToId = {}
		arn.scan('Anime', anime => {
			animeToId[anime.title.romaji] = anime.id
		}, () => {
			animeToIdJSONString = JSON.stringify(animeToId)

			response.render({
				user,
				animeToIdJSONString
			})
		})
	} else {
		response.render({
			user,
			animeToIdJSONString
		})
	}
}