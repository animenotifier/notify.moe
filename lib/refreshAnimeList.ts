import * as arn from './'
import * as Promise from 'bluebird'
import { User } from './interfaces/User'

// refreshAnimeList
export let refreshAnimeList = Promise.promisify(function(user: User, listProvider, animeProvider, airingDateProvider, listProviderSettings, cacheKey, callback) {
	return listProvider.getAnimeList(listProviderSettings.userName, (error, watchingOnProvider) => {
		if(error) {
			callback(error, watchingOnProvider)
			return
		}

		let mapToNativeAnime = watchingOnProvider.map(watchingAnime => arn.db.get('Anime', watchingAnime.id).then(anime => {
			return {
				id: anime.id,
				title: anime.title,
				image: anime.image,
				episodes: watchingAnime.episodes,
				preferredTitle: anime.title[user.titleLanguage]
			}
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
					console.error(error)

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
				return arn.db.set('AnimeLists', user.id, animeList).then(() => {
					callback(undefined, animeList)
				})
			}).catch(error => {
				callback(error, null)
			})
		})
	})
})