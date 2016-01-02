'use strict'

let request = require('request-promise')
let Promise = require('bluebird')
let xml2js = require('xml2js')
let specialTitles = require('./special.json')

function pad(pad, str, padLeft) {
	if(str === undefined)
		return pad

	if(padLeft)
		return (pad + str).slice(-pad.length)
	else
		return (str + pad).substring(0, pad.length)
}

class Nyaa {
	constructor() {
		this.headers = {
			'User-Agent': 'Anime Release Notifier'
		}

		this.xmlParser = new xml2js.Parser({
			explicitArray: false,	// Don't put single nodes into an array
			ignoreAttrs: true,		// Ignore attributes and only create text nodes
			trim: true,
			normalize: true,
			explicitRoot: false
		})

		this.episodeRegEx = / - (\d{2,3}) [\(\[]/

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
		anime.animeProvider.nextEpisodeUrl = anime.animeProvider.url + '+' + pad('00', anime.episodes.next, true)

		return request({
			uri: anime.animeProvider.rssUrl,
			method: 'GET',
			headers: this.headers
		}).then(body => {
			return this.xmlParser.parseStringAsync(body).then(json => {
				// Get highest episode number
				let episodes = json.channel.item.map(item => {
					let match = this.episodeRegEx.exec(item.title)
					if(match !== null) {
						//console.log(parseInt(match[1]), '=>', item.title)
						return parseInt(match[1])
					}

					return 0
				})

				let highestEpisodeNumber = Math.max.apply(Math, episodes)
				anime.episodes.available = highestEpisodeNumber
			})
		}).catch(error => {
			console.error(error)
		})
	}
}

module.exports = new Nyaa()