const maxEntries = 100

let Promise = require('bluebird')

exports.get = (request, response) => {
	if(!arn.auth(request, response, 'editor'))
		return

	let user = request.user

	let providers = [
		'Nyaa'
	]

	let matches = providers.reduce((obj, provider) => {
		obj[provider] = []
		return obj
	}, {})

	let tasks = providers.map(provider => arn.filter('AnimeTo' + provider, record => !record.edited && record.episodes === 0).then(uneditedMatches => matches[provider] = uneditedMatches))
	let animeIdToWatching = {}

	Promise.all(tasks).then(() => {
		Promise.all(providers.map(provider => {
			let providerMatches = matches[provider]
			let keys = providerMatches.map(match => match.id)

			return arn.batchGet('Anime', keys).then(results => {
				animeIdToWatching = results.reduce((dict, anime) => {
					if(anime.watching)
						dict[anime.id] = anime.watching

					return dict
				}, {})
				
				providerMatches = providerMatches.filter(match => animeIdToWatching[match.id] > 0)

				providerMatches.sort((a, b) => {
					let aWatching = animeIdToWatching[a.id]
					let bWatching = animeIdToWatching[b.id]

					if(!aWatching)
						return 1

					if(!bWatching)
						return -1

					return aWatching > bWatching ? -1 : 1
				})
				
				if(providerMatches.length > maxEntries)
					providerMatches.length = maxEntries
				
				matches[provider] = providerMatches
			})
		})).then(() => {
			response.render({
				user,
				matches,
				animeIdToWatching
			})
		})
	})
}