import * as arn from '../'
import * as request from 'request-promise'

class HummingBird {
	static headers = {
		'User-Agent': 'Anime Release Notifier',
		'Accept': 'application/json',
		'X-Client-Id': arn.api.hummingbird.v2.id
		//'X-Mashape-Key': arn.api.hummingbird.v1.secret
	}

	getAnimeListUrl(userName) {
		return `https://hummingbird.me/users/${userName}/library`
	}

	getAnimeList(userName, callback) {
		return request({
			uri: `https://hummingbird.me/api/v1/users/${userName}/library?status=currently-watching`,
			method: 'GET',
			headers: HummingBird.headers
		}).then(body => {
			let json: any = {}

			try {
				json = JSON.parse(body)
			} catch(error) {
				callback(error, [])
				return
			}

			let watching = json.map(entry => {
				let anime = entry.anime

				let episodesWatched = entry.episodes_watched ? parseInt(entry.episodes_watched) : 0
				let nextEpisodeToWatch = episodesWatched + 1
				let episodesOffset = 0

				return {
					title: anime.title,
					image: anime.cover_image,
					url: anime.url,
					providerId: parseInt(anime.id),
					airingDate: null,
					episodes: {
						watched: episodesWatched,
						next: nextEpisodeToWatch,
						available: 0,
						max: anime.episode_count ? parseInt(anime.episode_count) : -1,
						offset: episodesOffset
					}
				}
			})

			let tasks = new Array<Promise<any>>()
			watching.forEach(anime => {
				tasks.push(arn.getAnimeIdBySimilarTitle(anime, 'HummingBird').then(match => {
					anime.id = match ? match.id : null
					anime.similarity = match ? match.similarity : 0
				}))
			})

			Promise.all(tasks).then(() => callback(undefined, watching))
		}).catch(error => {
			callback(error, [])
		})
	}

	getUserImage(userName) {
		return request({
			uri: `https://hummingbird.me/api/v1/users/${userName}`,
			method: 'GET',
			headers: HummingBird.headers
		}).then(body => {
			let image = JSON.parse(body).avatar

			if(!image)
				throw new Error('No avatar')

			return image
		})
	}
}

module.exports = new HummingBird()