'use strict'

let arn = require('../../../lib')

let sortAlgorithms = {
	airingDate: (a, b) => {
		return a.airingDate.timeStamp - b.airingDate.timeStamp
	},

	alphabetically: (a, b) => {
		let aLower = a.title.toLowerCase()
		let bLower = b.title.toLowerCase()

		if(aLower < bLower)
			return -1

		if(aLower > bLower)
			return 1

		return 0
	}
}

module.exports = {
	get: function(request, response) {
		let nick = request.params[0]

		if(!nick)
			return response.end('Username not specified')

		arn.getUserByNickAsync(nick)
		.then(user => {
			let listProviderName = user.providers.list
			let listProvider = arn.listProviders[listProviderName]
			let animeProviderName = user.providers.anime
			let animeProvider = arn.animeProviders[animeProviderName]
			let airingDateProvider = arn.airingDateProviders[user.providers.airingDate]
			let listProviderSettings = user.listProviders[listProviderName]

			if(!listProvider)
				throw 'Invalid list provider'

			if(!listProviderSettings || !listProviderSettings.userName)
				throw `${listProviderName} username has not been specified`

			listProvider.getAnimeList(listProviderSettings.userName, function(error, watching) {
				let asyncTasks = []

				watching.forEach(function(entry) {
					entry.animeProvider = {
						url: null,
						nextEpisodeUrl: null,
						videoUrl: null
					}

					if(listProvider === airingDateProvider && airingDateProvider.getAiringDateById)
						asyncTasks.push(airingDateProvider.getAiringDateById(entry, entry.providerId))
					else
						asyncTasks.push(airingDateProvider.getAiringDate(entry))

					if(animeProvider)
						asyncTasks.push(animeProvider.getAnimeInfo(entry))
				})

				Promise.all(asyncTasks)
				.then(() => {
					watching.sort(sortAlgorithms['alphabetically'])

					let json = {
						listProvider: listProviderName,
						listUrl: listProvider.getAnimeListUrl(listProviderSettings.userName),
						watching
					}

					response.json(json)
				}).catch(error => {
					console.error(error)
					response.writeHead(409)
					response.end(error.toString())
				})
			})
		}).catch(error => {
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