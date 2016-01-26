'use strict'

let request = require('request-promise')

class HummingBird {
	constructor() {
		this.headers = {
			'User-Agent': 'Anime Release Notifier',
			'Accept': 'application/json',
			'X-Client-Id': arn.apiKeys.hummingbird.v2.clientID
			//'X-Mashape-Key': arn.apiKeys.hummingbird.v1.clientSecret
		}
	}

	getAnimeListUrl(userName) {
		return `https://hummingbird.me/users/${userName}/library`
	}

	getAnimeList(userName, callback) {
		return request({
			uri: `https://hummingbird.me/api/v1/users/${userName}/library?status=currently-watching`,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			let json = {}

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

			let tasks = []
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
}

module.exports = new HummingBird()