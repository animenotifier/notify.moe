'use strict'

let arn = require('../../../lib')

module.exports = {
	get: function(request, response) {
		let nick = request.params[0]

		if(!nick)
			return response.end('Username not specified')

		arn.getUserByNickAsync(nick)
		.then(user => {
			let listProviderName = user.providers.list
			let listProvider = arn[listProviderName]
			let listProviderSettings = user.listProviders[listProviderName]

			if(!listProvider)
				throw 'Invalid list provider'

			if(!listProviderSettings || !listProviderSettings.userName)
				throw `${listProviderName} username has not been specified`

			listProvider.getAnimeList(listProviderSettings.userName, function(error, watching) {
				watching.forEach(function(entry) {
					entry.animeProvider = {
						url: null,
						nextEpisodeUrl: null,
						videoUrl: null
					}
				})

				let json = {
					listProvider: listProviderName,
					listUrl: listProvider.getAnimeListUrl(listProviderSettings.userName),
					watching
				}

				response.setHeader('Content-Type', 'application/json')
				response.end(JSON.stringify(json))
			})
		})
		.catch(error => {
			response.writeHead(409)

			if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND')
				response.end(`User '${nick}' not found`)
			else if(error.message)
				response.end(error.message)
			else
				response.end(error.toString())
		})
	}
}