'use strict'

let arn = require('../../lib')

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

	let animeCount = 0
	arn.scan('Anime', function(anime) {
		animeCount++
	}, function() {
		response.render({
			user,
			animeCount
		})
	})
}