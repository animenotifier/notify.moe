'use strict'

let request = require('request-promise')
let Promise = require('bluebird')
let xml2js = require('xml2js')
let NodeCache = require('node-cache')
let specialTitles = require('./special.json')
let watch = require('node-watch')

function pad(pad, str, padLeft) {
	if(str === undefined)
		return pad

	if(str.length >= pad.length)
		return str

	if(padLeft)
		return (pad + str).slice(-pad.length)

	return (str + pad).substring(0, pad.length)
}

class Nyaa {
	constructor() {
		this.headers = {
			'User-Agent': 'Anime Release Notifier'
		}

		this.xmlParser = new xml2js.Parser({
			explicitArray: false,
			ignoreAttrs: true,
			trim: true,
			normalize: true,
			explicitRoot: false,
			strict: true
		})

		this.episodeRegEx = /[ _]-[ _](\d{2,3})[ _][\(\[]/

		this.cache = new NodeCache({
			stdTTL: 20 * 60
		})

		Promise.promisifyAll(this.cache)
		Promise.promisifyAll(this.xmlParser)

		this.getAnimeInfo = Promise.coroutine(function*(anime) {
			let searchTitle = yield arn.get('AnimeToNyaa', anime.id).then(match => {
				return match.title
			}).catch(error => {
				let tmpTitle = anime.title

				// Look up special.json
				if(specialTitles[tmpTitle]) {
					tmpTitle = specialTitles[tmpTitle]
				} else {
					tmpTitle = this.buildNyaaTitle(tmpTitle)
				}

				// Save in database
				if(tmpTitle) {
					arn.set('AnimeToNyaa', anime.id, {
						id: anime.id,
						title: tmpTitle
					})
				}

				return tmpTitle
			})

			searchTitle = searchTitle.replace(/ /g, '+')

			let quality = ''
			let subs = ''
			let nyaaSuffix = `&cats=1_37&filter=0&sort=2&term=${searchTitle}+${quality}+${subs}`.replace(/\+\+/g, '+').replace(/^\++|\++$/g, '')

			let nyaa = {
				url: `https://www.nyaa.se/?page=search${nyaaSuffix}`,
				rssUrl: `https://www.nyaa.se/?page=rss${nyaaSuffix}`,
				available: 0,
				type: 'download'
			}

			nyaa.nextEpisode = {
				url: nyaa.url + '+' + pad('00', anime.episodes.next.toString(), true)
			}

			let cacheKey = `${searchTitle}:${quality}:${subs}`

			yield this.cache.getAsync(cacheKey).then(available => {
				if(available) {
					nyaa.available = available
					return nyaa
				}

				return request({
					uri: nyaa.rssUrl,
					method: 'GET',
					headers: this.headers
				}).then(rss => {
					return this.getAvailableEpisodeCount(rss, anime, cacheKey).then(available => {
						nyaa.available = available
					})
				}).catch(error => {
					console.error(error, error.stack)
				})
			}).catch(error => {
				console.error(error, error.stack)
			})

			return nyaa
		})
	}

	buildNyaaTitle(title) {
		if(!title)
			return title

		title = title.replace(/[^[:alnum:]!']/gi, ' ')
		title = title.replace(/ \(?TV\)?/g, '')
		title = title.replace(/  /g, ' ')
		title = title.trim()
		return title
	}

	getAvailableEpisodeCount(rssResponse, anime, cacheKey) {
		return this.xmlParser.parseStringAsync(rssResponse).then(json => {
			let highestEpisodeNumber = 0

			if(Array.isArray(json.channel.item)) {
				// Get highest episode number
				let episodes = json.channel.item.map(item => {
					let match = this.episodeRegEx.exec(item.title)

					if(match !== null)
						return parseInt(match[1])

					return 0
				})

				highestEpisodeNumber = Math.max.apply(Math, episodes)
			}

			this.cache.set(cacheKey, highestEpisodeNumber, (error, success) => error)

			// Save available count in database
			arn.set('AnimeToNyaa', anime.id, {
				episodes: highestEpisodeNumber
			})

			return highestEpisodeNumber
		}).catch(error => {
			console.error(error, error.stack)
			return Promise.resolve(0)
		})
	}
}

module.exports = new Nyaa()