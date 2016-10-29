import * as arn from './'
import * as Promise from 'bluebird'
import { User } from './interfaces/User'

export const animeListCacheTime = 20 * 60 * 1000

export let getAnimeList = Promise.promisify(function getAnimeList(user: User, clearCache: boolean, callback) {
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

	let refresh = oldAnimeList => {
		let sendNotifications = animeList => {
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
		}

		return arn.refreshAnimeList(user, listProvider, animeProvider, airingDateProvider, listProviderSettings, cacheKey)
			.then(sendNotifications)
			.then(animeList => callback(undefined, animeList))
			.catch(callback)
	}

	arn.db.get('AnimeLists', user.id).then(animeList => {
		let now = new Date()
		let generated = new Date(animeList.generated)

		if(!clearCache && cacheKey === animeList.cacheKey && now.getTime() - generated.getTime() < arn.animeListCacheTime) {
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