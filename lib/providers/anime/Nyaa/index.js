'use strict'

let request = require('request-promise')

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
	}

	getAnimeInfo(anime) {
		let searchTitle = anime.title

		searchTitle = searchTitle.replace(/[^[:alnum:]!']/gi, ' ')
		searchTitle = searchTitle.replace(/ TV/g, '')
		searchTitle = searchTitle.replace(/  /g, ' ')
		searchTitle = searchTitle.trim()

		// TODO: Look up special.json

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

		}).catch(error => {
			console.error(error)
		})
	}
}

module.exports = new Nyaa()