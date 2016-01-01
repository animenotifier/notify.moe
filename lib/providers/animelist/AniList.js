'use strict'

let request = require('request')

class AniList {
	constructor() {
		this.authURL = 'https://anilist.co/api/auth/access_token'
		this.accessToken = undefined
		this.icon = 'http://anilist.co/favicon.png'
		this.cache = {}

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
		if(this.cache[userName]) {
			callback(undefined, this.cache[userName])
			return
		}

		if(!this.accessToken) {
			callback('No access token', [])
			return
		}

		let data = {}
		let apiURL = `https://anilist.co/api/user/${userName}/animelist?access_token=${this.accessToken}`

		request({
			uri: apiURL,
			method: 'GET',
			headers: {
				'User-Agent': 'Anime Release Notifier'
			}
		}, (error, response, body) => {
			let anilistJSON = {}

			try {
				anilistJSON = JSON.parse(body)
			} catch(e) {
				callback(e, [])
				return
			}

			let lists = anilistJSON.lists

			if(!lists.watching) {
				callback('Your anime list doesn\'t include a watching list.', [])
				return
			}

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

			this.cache[userName] = watching
			callback(error, watching)
		})
	}
}

module.exports = new AniList()