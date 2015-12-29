'use strict'

let request = require('request')

class AniList {
	constructor() {
		this.authURL = 'https://anilist.co/api/auth/access_token'
		this.accessToken = undefined

		// Authorize every 30 minutes
		setInterval(this.authorize.bind(this), 30 * 60 * 1000)

		// Authorize now
		this.authorize()
	}

	authorize() {
		request({
			uri: this.authURL,
			method: 'POST',
			json: {
				grant_type: 'client_credentials',
				client_id: 'akyoto-wbdln',
				client_secret: 'zS3MidMPmolyHRYNOvSR1'
			},
			headers: {
				'User-Agent': 'Anime Release Notifier'
			}
		}, (error, response, body) => {
			this.accessToken = body.access_token
			console.log('Successfully authorized AniList API access!')
		})
	}

	getAnimeListUrl(userName) {
		return `https://anilist.co/animelist/${userName}`
	}

	getAnimeList(userName, callback) {
		let data = {}
		let apiURL = `https://anilist.co/api/user/${userName}/animelist?access_token=${this.accessToken}`

		request({
			uri: apiURL,
			method: 'GET',
			headers: {
				'User-Agent': 'Anime Release Notifier'
			}
		}, (error, response, body) => {
			let anilistJSON = JSON.parse(body)
			let lists = anilistJSON.lists
			let watching = []

			lists.watching.forEach(entry => {
				let anime = entry.anime

				let episodesWatched = entry.episodes_watched
				let nextEpisodeToWatch = episodesWatched + 1
				let episodesOffset = 0

				let newEntry = {
					title: anime.title_english.trim(),
					image: anime.image_url_lge.replace('http://', 'https://'),
					url: 'https://anilist.co/anime/' + anime.id,
					providerId: anime.id,
					airingDate: {
						timeStamp: null,
						remaining: null,
						remainingString: ''
					},
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

			callback(error, watching)
		})
	}
}

module.exports = new AniList()