import * as arn from '../'
import * as querystring from 'querystring'
import * as request from 'request-promise'

class AniList {
	static authURL = 'https://anilist.co/api/auth/access_token'
	static icon = 'http://anilist.co/favicon.png'
	static headers = {
		'User-Agent': 'Anime Release Notifier',
		'Accept': 'application/json'
	}

	accessToken = undefined
	security: any

	constructor() {
		this.security = require('../../security/api-keys.json').anilist
		this.authorize()

		// Authorize every 30 minutes
		setInterval(this.authorize.bind(this), 30 * 60 * 1000)
	}

	authorize() {
		return request({
			uri: AniList.authURL,
			method: 'POST',
			json: {
				grant_type: 'client_credentials',
				client_id: this.security.id,
				client_secret: this.security.secret
			},
			headers: AniList.headers
		})
		.then(body => this.accessToken = body.access_token)
		.catch(error => console.error('Anilist access token acquisition failed'))
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
			headers: AniList.headers
		}).then(body => {
			let anilistJSON: any = {}

			try {
				anilistJSON = JSON.parse(body)
			} catch(error) {
				callback(error, [])
				return
			}

			let lists = anilistJSON.lists

			if(lists.watching === undefined)
				lists.watching = []

			let watching = lists.watching.map(entry => {
				let anime = entry.anime

				if(!anime)
					return null

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
			}).filter(entry => entry !== null)

			callback(undefined, watching)
		}).catch(error => {
			console.error(apiURL, error)
			callback(error, [])
		})
	}

	getAiringDate(anime) {
		let now = (new Date()).getTime() / 1000

		return request({
			uri: `https://anilist.co/api/anime/${anime.id}/airing?access_token=${this.accessToken}`,
			method: 'GET',
			headers: AniList.headers
		}).then(body => {
			let anilistJSON = JSON.parse(body)
			let timeStamp: number = anilistJSON[anime.episodes.next]

			// Airing date not available for this episode?
			if(!timeStamp)
				throw new Error(`No airing date available for ${anime.id}`)

			let remaining = timeStamp - now
			let remainingString = remaining + arn.plural(remaining, 'second')

			let days = arn.inDays(now, timeStamp)
			if(Math.abs(days) >= 1) {
				remainingString = arn.plural(days, 'day')
			} else {
				let hours = arn.inHours(now, timeStamp)
				if(Math.abs(hours) >= 1) {
					remainingString = arn.plural(hours, 'hour')
				} else {
					let minutes = arn.inMinutes(now, timeStamp)
					if(Math.abs(minutes) >= 1) {
						remainingString = arn.plural(minutes, 'minute')
					} else {
						let seconds = arn.inSeconds(now, timeStamp)
						remainingString = arn.plural(seconds, 'second')
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
			if(!error.message || !error.message.startsWith('No airing date available'))
				console.error(error)

			return {
				timeStamp: null,
				remaining: null,
				remainingString: ''
			}
		})
	}

	getAnimeFromPage(page) {
		return request({
			uri: `https://anilist.co/api/browse/anime?sort=updated_at-desc&page=${page}&access_token=${this.accessToken}`,
			method: 'GET',
			headers: AniList.headers
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
					airingStatus: anime.airing_status ? anime.airing_status : '',
					adult: anime.adult ? 1 : 0,
					anilistEdited: anime.updated_at
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
			headers: AniList.headers
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
			if(error.name === 'StatusCodeError') {
				console.warn(`Unavailable [${error.statusCode}]: ${error.options.uri}`)
				return
			}

			console.error(`Error importing anime details for ID ${animeId}`, error.stack)
		})
	}

	getUserImage(userName) {
		return request({
			uri: `https://anilist.co/api/user/${userName}?access_token=${this.accessToken}`,
			method: 'GET',
			headers: AniList.headers
		}).then(body => {
			let image = JSON.parse(body).image_url_lge

			if(image.endsWith('default.png'))
				throw new Error('Default avatar')

			return image.replace('http://', 'https://')
		})
	}
}

module.exports = new AniList()