import * as arn from 'arn'
import { User } from 'arn/interfaces/User'

export const animeListCacheTime = 20 * 60 * 1000

async function sendNotifications(user: User, animeList, oldAnimeList) {
	if(!oldAnimeList)
		return animeList

	// Did the user enable notifications?
	if(Object.keys(user.pushEndpoints).length === 0)
		return animeList

	// Compare to check if we can send notifications
	await animeList.watching.map(anime => {
		let oldAnime = oldAnimeList.watching.find(e => e.id === anime.id)

		if(!oldAnime)
			return

		let shouldSendNotification = (
			anime.episodes &&
			oldAnime.episodes &&
			anime.episodes.available === anime.episodes.next &&
			anime.episodes.available === oldAnime.episodes.available + 1
		)

		if(shouldSendNotification) {
			// Send push notification to the user
			return arn.sendNotification(user, {
				title: anime.preferredTitle,
				icon: anime.image,
				body: `Episode ${anime.episodes.available} was just released`
			})
		}

		return Promise.resolve()
	})

	return animeList
}

export function getAnimeList(user: User, clearCache: boolean): Promise<any> {
	const listProviderName = user.providers.list
	const listProvider = arn.listProviders[listProviderName]
	const animeProviderName = user.providers.anime
	const animeProvider = arn.animeProviders[animeProviderName]
	const airingDateProvider = arn.airingDateProviders[user.providers.airingDate]
	const listProviderSettings = user.listProviders[listProviderName]

	if(!listProvider)
		throw new Error('Invalid list provider')

	if(!listProviderSettings || !listProviderSettings.userName)
		throw new Error(`${listProviderName} username has not been specified`)

	let cacheKey = listProviderName + ':' + listProviderSettings.userName + ':' + animeProviderName + ':' + user.sortBy + ':' + user.titleLanguage

	let refresh = function(oldAnimeList: any | undefined) {
		return arn.refreshAnimeList(user, listProvider, animeProvider, airingDateProvider, listProviderSettings, cacheKey)
		.then(animeList => sendNotifications(user, animeList, oldAnimeList))
	}

	return arn.db.get('AnimeLists', user.id).then(animeList => {
		let now = new Date()
		let generated = new Date(animeList.generated)

		if(arn.production && !clearCache && cacheKey === animeList.cacheKey && now.getTime() - generated.getTime() < arn.animeListCacheTime) {
			return animeList
		} else {
			return refresh(animeList)
		}
	}).catch(error => {
		if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND')
			return refresh(undefined)

		throw error
	})
}