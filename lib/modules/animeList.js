'use strict'

let Promise = require('bluebird')

// getAnimeList
arn.getAnimeList = Promise.promisify(function(user, clearCache, callback) {
	let listProviderName = user.providers.list
	let listProvider = arn.listProviders[listProviderName]
	let animeProviderName = user.providers.anime
	let animeProvider = arn.animeProviders[animeProviderName]
	let airingDateProvider = arn.airingDateProviders[user.providers.airingDate]
	let listProviderSettings = user.listProviders[listProviderName]

	if(!listProvider)
		callback(new Error('Invalid list provider'))

	if(!listProviderSettings || !listProviderSettings.userName)
		callback(new Error(`${listProviderName} username has not been specified`))

	let cacheKey = listProviderName + ':' + listProviderSettings.userName + ':' + animeProviderName + ':' + user.sortBy + ':' + user.titleLanguage

	let refresh = (oldAnimeList) => {
		return arn.refreshAnimeList(user, listProvider, animeProvider, airingDateProvider, listProviderSettings, cacheKey).then(animeList => {
			if(!oldAnimeList)
				return animeList

			// Did the user enable notifications?
			if(Object.keys(user.pushEndpoints).length === 0)
				return animeList

			// Compare to check if we can send notifications
			animeList.watching.forEach(anime => {
				let oldAnime = oldAnimeList.watching.find(e => e.id === anime.id)

				if(!oldAnime)
					return

				// Send push notification
				if(
					anime.episodes &&
					oldAnime.episodes &&
					anime.episodes.available === anime.episodes.next &&
					anime.episodes.available === oldAnime.episodes.available + 1
				) {
					return arn.sendNotification(user, {
						title: anime.preferredTitle,
						icon: anime.image,
						body: `Episode ${anime.episodes.available} was just released`
					})
				}
			})

			return animeList
		})
		.then(animeList => callback(undefined, animeList))
		.catch(callback)
	}

	arn.get('AnimeLists', user.id).then(animeList => {
		let now = new Date()
		let generated = new Date(animeList.generated)

		if(arn.cacheAnimeLists && !clearCache && cacheKey === animeList.cacheKey && now.getTime() - generated.getTime() < arn.animeListCacheTime) {
			callback(undefined, animeList)
		} else {
			return refresh(animeList)
		}
	}).catch(error => {
		if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND') {
			return refresh()
		} else {
			callback(error)
		}
	})
})

// refreshAnimeList
arn.refreshAnimeList = Promise.promisify(function(user, listProvider, animeProvider, airingDateProvider, listProviderSettings, cacheKey, callback) {
	return listProvider.getAnimeList(listProviderSettings.userName, (error, watchingOnProvider) => {
		if(error) {
			callback(error, watchingOnProvider)
			return
		}

		let mapToNativeAnime = watchingOnProvider.map(watchingAnime => arn.get('Anime', watchingAnime.id).then(anime => {
			anime.episodes = watchingAnime.episodes
			anime.preferredTitle = anime.title[user.titleLanguage]
			return anime
		}))

		Promise.all(mapToNativeAnime).then(watching => {
			let tasks = []

			watching.forEach(entry => {
				// Airing date
				tasks.push(airingDateProvider.getAiringDate(entry).then(airingDate => entry.airingDate = airingDate))

				// Anime provider
				tasks.push(animeProvider.getAnimeInfo(entry).then(animeInfo => {
					entry.animeProvider = animeInfo
					entry.episodes.available = entry.animeProvider.available
				}).catch(error => {
					console.error(error, error.stack)

					entry.animeProvider = {
						url: null,
						nextEpisode: null,
						available: 0
					}
					entry.episodes.available = 0
				}))
			})

			Promise.all(tasks).then(() => {
				watching.sort(arn.sortAlgorithms[user.sortBy])

				let animeList = {
					user: user.nick,
					userId: user.id,
					listProvider: user.providers.list,
					listUrl: listProvider.getAnimeListUrl(listProviderSettings.userName),
					titleLanguage: user.titleLanguage,
					watching,
					cacheKey,
					generated: (new Date()).toISOString()
				}

				// Cache it
				return arn.set('AnimeLists', user.id, animeList).then(() => {
					callback(undefined, animeList)
				})
			}).catch(error => {
				callback(error, null)
			})
		})
	})
})

// getAnimeListByNick
arn.getAnimeListByNick = function(nick, clearCache) {
	return arn.getUserByNick(nick).then(user => arn.getAnimeList(user, clearCache)).catch(error => {
		console.error(error, error.stack)

		if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND')
			return Promise.reject(`User '${nick}' not found`)

		if(error.message)
			return Promise.reject(error.message)

		return Promise.reject(error.toString())
	})
}