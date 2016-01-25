'use strict'

let request = require('request-promise')
let Promise = require('bluebird')
let fs = Promise.promisifyAll(require('fs'))
let xml2js = Promise.promisifyAll(require('xml2js'))
let striptags = require('striptags')

class CrunchyRoll {
	constructor() {
		this.headers = {
			'User-Agent': 'Anime Release Notifier'
		}

		this.cacheTime = 20 * 60 * 1000

		this.rssLinkRegEx = /(http:\/\/www\.crunchyroll\.com\/[^"]+?\.rss)/

		this.xmlParser = new xml2js.Parser({
			explicitArray: true,	// Always put child nodes in an array if true; otherwise an array is created only if there is more than one.
			ignoreAttrs: true,		// Ignore all XML attributes and only create text nodes.
			trim: true,				// Trim the whitespace at the beginning and end of text nodes
			normalize: true,		// Trim whitespaces inside text nodes.
			explicitRoot: false,	// Set this if you want to get the root node in the resulting object.
			strict: true			// Set sax-js to strict or non-strict parsing mode. Defaults to true which is highly recommended.
		})

		this.getAnimeInfo = Promise.coroutine(function*(anime) {
			let crunchy = yield arn.get('CrunchyRoll', anime.id).catch(error => {
				return {
					url: null,
					rssUrl: null,
					available: 0
				}
			})

			let crunchyLink = anime.links.find(link => link.url.indexOf('crunchyroll.com') !== -1)

			if(!crunchyLink || !crunchyLink.url)
				return crunchy

			crunchy.url = crunchyLink.url

			// Get RSS URL
			if(!crunchy.rssUrl) {
				//let response = yield fs.readFileAsync('security/crunchyroll-response.html', 'utf8')
				let response = yield request({
					uri: crunchy.url,
					method: 'GET',
					headers: this.headers
				})

				let match = this.rssLinkRegEx.exec(response)

				if(match === null) {
					console.error('No crunchyroll RSS link found')
					return crunchy
				}

				crunchy.rssUrl = match[1]
			}

			// Get episode list
			if(!crunchy.generated || (new Date()).getTime() - (new Date(crunchy.generated)).getTime() <= this.cacheTime) {
				//let xmlResponse = yield fs.readFileAsync('security/crunchyroll.xml', 'utf8')
				let xmlResponse = yield request({
					uri: crunchy.rssUrl,
					method: 'GET',
					headers: this.headers
				})

				let rss = yield this.xmlParser.parseStringAsync(xmlResponse)

				let episodes = rss.channel[0].item

				if(!episodes)
					return crunchy

				crunchy.episodes = episodes.map(episode => {
					let pubDate = new Date(episode.pubDate[0]).getTime()

					// Do not include episodes whose publishing date is in the future
					if(pubDate - (new Date()).getTime() > 0)
						return null

					return {
						number: parseInt(episode['crunchyroll:episodeNumber'][0]),
						title: episode.title[0],
						url: episode.link[0],
						description: striptags(episode.description[0]),
						timeStamp: pubDate
					}
				}).filter(episode => episode !== null)

				crunchy.id = anime.id
				crunchy.available = crunchy.episodes.length
				crunchy.generated = (new Date()).toISOString()
			}

			// Cache it
			arn.set('CrunchyRoll', anime.id, crunchy)

			// User-specific: next episode with URL in it
			let nextEpisode = crunchy.episodes.find(episode => episode.number === anime.episodes.next)

			if(nextEpisode)
				crunchy.nextEpisode = nextEpisode
			else
				crunchy.nextEpisode = null

			// We don't need those in the user anime list
			delete crunchy.id
			delete crunchy.episodes
			crunchy.type = 'watch'

			return crunchy
		})
	}
}

module.exports = new CrunchyRoll()