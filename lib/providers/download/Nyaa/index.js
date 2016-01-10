'use strict'

let request = require('request-promise')
let Promise = require('bluebird')
let xml2js = require('xml2js')
let NodeCache = require('node-cache')
let specialTitles = require('./special.json')

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

		this.episodeRegEx = / - (\d{2,3}) [\(\[]/

		this.cache = new NodeCache({
			stdTTL: 20 * 60
		})

		Promise.promisifyAll(this.cache)
		Promise.promisifyAll(this.xmlParser)
	}

	getAnimeInfo(anime) {
		let searchTitle = anime.title

		searchTitle = searchTitle.replace(/[^[:alnum:]!']/gi, ' ')
		searchTitle = searchTitle.replace(/ \(?TV\)?/g, '')
		searchTitle = searchTitle.replace(/  /g, ' ')
		searchTitle = searchTitle.trim()

		// Look up special.json
		if(specialTitles[searchTitle])
			searchTitle = specialTitles[searchTitle]

		searchTitle = searchTitle.replace(/ /g, '+')

		let quality = ''
		let subs = ''
		let nyaaSuffix = `&cats=1_37&filter=0&sort=2&term=${searchTitle}+${quality}+${subs}`.replace(/\+\+/g, '+').replace(/^\++|\++$/g, '')

		anime.animeProvider.url = `http://www.nyaa.se/?page=search${nyaaSuffix}`
		anime.animeProvider.rssUrl = `http://www.nyaa.se/?page=rss${nyaaSuffix}`
		anime.animeProvider.nextEpisodeUrl = anime.animeProvider.url + '+' + pad('00', anime.episodes.next.toString(), true)

		let cacheKey = `${searchTitle}:${quality}:${subs}`
		return this.cache.getAsync(cacheKey).then(available => {
			if(available) {
				anime.episodes.available = available
				return Promise.resolve(anime.episodes.available)
			}

			return request({
				uri: anime.animeProvider.rssUrl,
				method: 'GET',
				headers: this.headers
			}).then(body => {
				return this.xmlParser.parseStringAsync(body).then(json => {
					let highestEpisodeNumber = 0

					if(Array.isArray(json.channel.item)) {
						// Get highest episode number
						let episodes = json.channel.item.map(item => {
							let match = this.episodeRegEx.exec(item.title)
							if(match !== null) {
								//console.log(parseInt(match[1]), '=>', item.title)
								return parseInt(match[1])
							}

							return 0
						})

						highestEpisodeNumber = Math.max.apply(Math, episodes)
					}

					anime.episodes.available = highestEpisodeNumber
					this.cache.set(cacheKey, highestEpisodeNumber, (error, success) => error)
				})
			}).catch(error => {
				console.error(error.stack)
			})
		}).catch(error => {
			console.error(error.stack)
		})
	}
}

module.exports = new Nyaa()