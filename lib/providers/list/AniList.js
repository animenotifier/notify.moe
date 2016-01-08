'use strict'

let request = require('request-promise')
let plural = require('../../utils/plural')
let datediff = require('../../utils/datediff')
let apiKeys = require('../../../security/api-keys.json')
let Promise = require('bluebird')
let querystring = require('querystring')

class AniList {
	constructor() {
		this.authURL = 'https://anilist.co/api/auth/access_token'
		this.accessToken = undefined
		this.icon = 'http://anilist.co/favicon.png'
		this.lastReplyCount = undefined
		this.headers = {
			'User-Agent': 'Anime Release Notifier',
			'Accept': 'application/json'
		}

		// Authorize every 30 minutes
		setInterval(this.authorize.bind(this), 1 * 5 * 1000)

		// Authorize now
		this.authorize().then(() => {
			console.log('Successfully authorized AniList API access!')

			// Check forum thread replies
			setInterval(this.checkForumReplies.bind(this), 5 * 60 * 1000)
			this.checkForumReplies()
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

	checkForumReplies() {
		let arn = require('../../../lib')
		let forumThreadId = 64
		let apiURL = `https://anilist.co/api/forum/thread/${forumThreadId}?access_token=${this.accessToken}`

		return request({
			uri: apiURL,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			let thread = JSON.parse(body)
			if(!this.lastReplyCount || thread.reply_count > this.lastReplyCount) {
				console.log(`Anilist forum thread has ${thread.reply_count} replies`)

				if(this.lastReplyCount !== undefined)
					arn.events.emit('new forum reply', `https://anilist.co/forum/thread/${forumThreadId}`, thread.reply_user.display_name)

				this.lastReplyCount = thread.reply_count
			}
		}).catch(error => {
			console.error('Error checking anilist.co forum replies:', error)
		})
	}

	getAnimeListUrl(userName) {
		return `https://anilist.co/animelist/${userName}`
	}

	getAnimeList(userName, callback) {
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

			callback(undefined, watching)
		}).catch(error => {
			console.error(error.stack)
			callback(error, [])
		})
	}

	getAiringDate(anime) {
		let searchTitle = querystring.stringify({a: anime.title.replace('/', ' ')}).substring(2)

		return request({
			uri: `https://anilist.co/api/anime/search/${searchTitle}?access_token=${this.accessToken}`,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			body = body.trim()
			let json = null

			try {
				json = JSON.parse(body)
			} catch(error) {
				throw new Error(`Couldn\'t get airing date for ${anime.title} (${searchTitle})`)
			}

			let id = json[0].id

			if(id === undefined)
				throw new Error(`Undefined anilist ID for anime '${json[0].title_english}'`)

			return id
		})
		.then(id => this.getAiringDateById(anime, id))
		.catch(error => {
			if(!anime.airingDate) {
				anime.airingDate = {
					timeStamp: null,
					remaining: null,
					remainingString: ''
				}
			}

			if(error.statusCode !== 404)
				console.error(error.stack)
		})
	}

	getAiringDateById(anime, anilistId) {
		anime.airingDate = {
			timeStamp: null,
			remaining: null,
			remainingString: ''
		}

		if(!anilistId)
			return Promise.resolve()

		let now = (new Date()).getTime() / 1000

		return request({
			uri: `https://anilist.co/api/anime/${anilistId}/airing?access_token=${this.accessToken}`,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			let anilistJSON = JSON.parse(body)
			let timeStamp = anilistJSON[anime.episodes.next]

			// Airing date not available for this episode?
			if(!timeStamp)
				return

			let remaining = parseInt(timeStamp - now)
			let remainingString = remaining + plural(remaining, 'second')

			let days = datediff.inDays(now, timeStamp)
			if(days !== 0) {
				remainingString = plural(days, 'day')
			} else {
				let hours = datediff.inHours(now, timeStamp)
				if(hours !== 0) {
					remainingString = plural(hours, 'hour')
				} else {
					let minutes = datediff.inMinutes(now, timeStamp)
					remainingString = plural(minutes, 'minute')
				}
			}

			// Add 'ago' if the date is in the past
			if(remainingString.startsWith('-')) {
				remainingString = remainingString.substring(1) + ' ago'
			}

			anime.airingDate = {
				timeStamp: timeStamp,
				remaining: remaining,
				remainingString: remainingString
			}
		}).catch(error => {
			console.error(error.stack)
		})
	}
}

module.exports = new AniList()