'use strict'

let arn = require('../../../lib')

module.exports = {
	get: function(request, response) {
		let nick = request.params[0]

		if(!nick)
			return response.end()

		let anilistNick = 'Akyoto'
		
		arn.AniList.getAnimeList(anilistNick, function(error, data) {
			response.setHeader('Content-Type', 'application/json')
			response.end(JSON.stringify(data))
		})
	}
}