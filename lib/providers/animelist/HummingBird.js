'use strict'

let request = require('request-promise')
let plural = require('../../plural')
let datediff = require('../../datediff')
let apiKeys = require('../../../security/api-keys.json')

class HummingBird {
	constructor() {
		this.headers = {
			'User-Agent': 'Anime Release Notifier',
			'Accept': 'application/json',
			//'X-Client-Id': apiKeys.hummingbird.clientID,
			'X-Mashape-Key': apiKeys.hummingbird.v1.clientSecret
		}
	}

	getAnimeListUrl(userName) {
		return `https://hummingbird.me/users/${userName}/library`
	}

	getAnimeList(userName, callback) {
		let apiURL = `https://hummingbirdv1.p.mashape.com/users/${userName}/library?status=currently-watching`

		request({
			uri: apiURL,
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

			let watching = []

			json.forEach(entry => {
				let anime = entry.anime

				let episodesWatched = entry.episodes_watched
				let nextEpisodeToWatch = episodesWatched + 1
				let episodesOffset = 0

				let newEntry = {
					title: anime.title,
					image: anime.cover_image,
					url: anime.url,
					providerId: anime.id,
					airingDate: null,
					episodes: {
						watched: episodesWatched ? episodesWatched : 0,
						next: nextEpisodeToWatch,
						available: 0,
						max: anime.total_episodes ? anime.total_episodes : -1,
						offset: episodesOffset,
					}
				}

				watching.push(newEntry)
			})

			callback(undefined, watching)
		}).catch(error => {
			callback(error, [])
		})
	}
}

module.exports = new HummingBird()