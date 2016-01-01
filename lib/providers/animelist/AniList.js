'use strict'

let request = require('request-promise')
let plural = require('../../plural')
let datediff = require('../../datediff')
let apiKeys = require('../../../security/api-keys.json')
let Promise = require('bluebird')

class AniList {
	constructor() {
		this.authURL = 'https://anilist.co/api/auth/access_token'
		this.accessToken = undefined
		this.icon = 'http://anilist.co/favicon.png'
		this.headers = {
			'User-Agent': 'Anime Release Notifier',
			'Accept': 'application/json'
		}
		//this.cache = {}

		// Authorize every 30 minutes
		setInterval(this.authorize.bind(this), 30 * 60 * 1000)

		// Authorize now
		this.authorize().then(() => {
			console.log('Successfully authorized AniList API access!')
		})
	}

	authorize() {
		return request({
			uri: this.authURL,
			method: 'POST',
			json: {
				grant_type: 'client_credentials',
				client_id: apiKeys.anilist.clientID,
				client_secret: apiKeys.anilist.clientSecret
			},
			headers: this.headers
		}).then(body => this.accessToken = body.access_token)
	}

	getAnimeListUrl(userName) {
		return `https://anilist.co/animelist/${userName}`
	}

	getAnimeList(userName, callback) {
		if(this.cache && this.cache[userName]) {
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
			headers: this.headers
		}).then(body => {
			let anilistJSON = {}

			try {
				anilistJSON = JSON.parse(body)
			} catch(error) {
				callback(error, [])
				return
			}

			let lists = anilistJSON.lists

			if(!lists.watching) {
				callback('Your anime list doesn\'t include a watching list.', [])
				return
			}

			let watching = lists.watching.map(entry => {
				let anime = entry.anime

				let episodesWatched = entry.episodes_watched ? parseInt(entry.episodes_watched) : 0
				let nextEpisodeToWatch = episodesWatched + 1
				let episodesOffset = 0

				return {
					title: anime.title_english.trim(),
					image: anime.image_url_lge.replace('http://', 'https://'),
					url: 'https://anilist.co/anime/' + anime.id,
					providerId: parseInt(anime.id),
					airingDate: null,
					episodes: {
						watched: episodesWatched,
						next: nextEpisodeToWatch,
						available: 0,
						max: anime.total_episodes ? anime.total_episodes : -1,
						offset: episodesOffset
					}
				}
			})

			if(this.cache)
				this.cache[userName] = watching

			callback(undefined, watching)
		}).catch(error => {
			callback(error, watching)
		})
	}

	getAiringDate(anime) {
		return request({
			uri: `https://anilist.co/api/anime/search/${anime.title}?access_token=${this.accessToken}`,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			anime.airingDate = {
				timeStamp: null,
				remaining: null,
				remainingString: ''
			}

			body = body.trim()
			if(!body)
				return Promise.resolve()

			let json = null

			try {
				json = JSON.parse(body)
			} catch(error) {
				console.error(error)
				return Promise.resolve()
			}

			if(json.length === 0)
				return Promise.resolve()

			return this.getAiringDateById(anime, json[0].id)
		})
	}

	getAiringDateById(anime, anilistId) {
		let now = (new Date()).getTime() / 1000

		return request({
			uri: `https://anilist.co/api/anime/${anilistId}/airing?access_token=${this.accessToken}`,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			let anilistJSON = {}

			try {
				anilistJSON = JSON.parse(body)

				let timeStamp = anilistJSON[anime.episodes.next]

				if(!timeStamp)
					throw 'Airing date not available for episode ' + anime.episodes.next

				let remaining = parseInt(timeStamp - now)
				let remainingString = remaining + plural(remaining, 'second')

				let days = datediff.inDays(now, timeStamp)
				if(days !== 0) {
					remainingString = plural(days, 'day')
				} else {
					let hours = datediff.inHours(now, timeStamp)
					remainingString = plural(hours, 'hour')
				}

				anime.airingDate = {
					timeStamp: timeStamp,
					remaining: remaining,
					remainingString: remainingString
				}
			} catch(error) {
				console.log(error)
				anime.airingDate = {
					timeStamp: null,
					remaining: null,
					remainingString: ''
				}
			}
		})
	}
}

module.exports = new AniList()