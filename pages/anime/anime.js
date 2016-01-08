'use strict'

let arn = require('../../lib')

exports.get = function(request, response) {
	let user = request.user
	let animeId = request.params[0]

	if(animeId) {
		arn.get('Anime', animeId).then(anime => {
			response.render({
				user,
				anime
			})
		}).catch(error => {
			response.render({
				user
			})
		})
		return
	}

	response.render({
		user
	})
}