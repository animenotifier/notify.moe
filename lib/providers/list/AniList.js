'use strict'

let request = require('request-promise')
let plural = require('../../utils/plural')
let datediff = require('../../utils/datediff')
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
	}

	authorize() {
		return request({
			uri: this.authURL,
			method: 'POST',
			json: {
				grant_type: 'client_credentials',
				client_id: arn.apiKeys.anilist.clientID,
				client_secret: arn.apiKeys.anilist.clientSecret
			},
			headers: this.headers
		}).then(body => this.accessToken = body.access_token)
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
					id: parseInt(anime.id),
					similarity: 1,
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
		let now = (new Date()).getTime() / 1000

		return request({
			uri: `https://anilist.co/api/anime/${anime.id}/airing?access_token=${this.accessToken}`,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			let anilistJSON = JSON.parse(body)
			let timeStamp = anilistJSON[anime.episodes.next]

			// Airing date not available for this episode?
			if(!timeStamp)
				throw new Error(`No airing date available for ${anime.id}`)

			let remaining = parseInt(timeStamp - now)
			let remainingString = remaining + plural(remaining, 'second')

			let days = datediff.inDays(now, timeStamp)
			if(Math.abs(days) >= 1) {
				remainingString = plural(days, 'day')
			} else {
				let hours = datediff.inHours(now, timeStamp)
				if(Math.abs(hours) >= 1) {
					remainingString = plural(hours, 'hour')
				} else {
					let minutes = datediff.inMinutes(now, timeStamp)
					if(Math.abs(minutes) >= 1) {
						remainingString = plural(minutes, 'minute')
					} else {
						let seconds = datediff.inSeconds(now, timeStamp)
						remainingString = plural(seconds, 'second')
					}
				}
			}

			// Add 'ago' if the date is in the past
			if(remainingString.startsWith('-')) {
				remainingString = remainingString.substring(1) + ' ago'
			}

			return {
				timeStamp: timeStamp,
				remaining: remaining,
				remainingString: remainingString
			}
		}).catch(error => {
			console.error(error.stack)

			return {
				timeStamp: null,
				remaining: null,
				remainingString: ''
			}
		})
	}

	getAnimeFromPage(page) {
		return request({
			uri: `https://anilist.co/api/browse/anime?sort=id&page=${page}&access_token=${this.accessToken}`,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			let anilist = JSON.parse(body)

			return anilist.map(anime => {
				return {
					id: parseInt(anime.id),
					type: anime.type,
					title: {
						romaji: anime.title_romaji,
						japanese: anime.title_japanese,
						english: anime.title_english,
						synonyms: anime.synonyms
					},
					image: anime.image_url_lge.replace('http://', 'https://'),
					airingStatus: anime.airing_status,
					adult: anime.adult ? 1 : 0
				}
			})
		}).catch(error => {
			console.error(`Error importing anime from page ${page}:`, error.stack)
		})
	}

	getAnimeDetails(animeId) {
		return request({
			uri: `https://anilist.co/api/anime/${animeId}/page?access_token=${this.accessToken}`,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			let anime = JSON.parse(body)

			return {
				description: (anime.description && anime.description !== null && anime.description !== 'N/A') ? anime.description : null,
				startDate: anime.start_date,
				endDate: anime.end_date,
				hashtag: anime.hashtag ? anime.hashtag : null,
				youtubeId: anime.youtube_id ? anime.youtube_id : null,
				genres: anime.genres ? anime.genres : [],
				source: anime.source ? anime.source : null,
				classification: anime.classification ? anime.classification : null,
				totalEpisodes: anime.total_episodes,
				duration: anime.duration,
				links: anime.external_links.map(link => {
					return {
						url: link.url,
						title: link.site ? link.site : ''
					}
				}),
				studios: anime.studio.map(studio => {
					return {
						id: studio.id,
						name: studio.studio_name,
						wiki: studio.studio_wiki ? studio.studio_wiki : null,
						isMainStudio: (studio.main_studio && studio.main_studio !== null && studio.main_studio !== 0) ? 1 : 0
					}
				}),
				relations: anime.relations.map(relation => {
					return {
						id: relation.id,
						type: relation.relation_type
					}
				})
			}
		}).catch(error => {
			console.error(`Error importing anime details for ID ${animeId}`, error.stack)
		})
	}

	checkForumReplies() {
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
}

module.exports = new AniList()