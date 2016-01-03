'use strict'

let request = require('request-promise')
let xml2js = require('xml2js')
let Promise = require('bluebird')
let arn = require('../lib')

module.exports = function(aero) {
	let xmlParser = new xml2js.Parser({
		explicitArray: false,	// Don't put single nodes into an array
		ignoreAttrs: true,		// Ignore attributes and only create text nodes
		trim: true,
		normalize: true,
		explicitRoot: false
	})

	Promise.promisifyAll(xmlParser)

	let updateNews = function() {
		let url = 'http://www.animenewsnetwork.com/news/rss.xml'

		request(url)
		.then(rss => {
			return xmlParser.parseStringAsync(rss).then(json => {
				let items = json.channel.item

				arn.news = {
					'Anime': items.filter(item => item.category === 'Anime'),
					'Manga': items.filter(item => item.category === 'Manga'),
					'Games': items.filter(item => item.category === 'Games')
				}
			})
		})
		.catch(error => {
			console.log(error)
		})
	}

	setInterval(updateNews, 15 * 60 * 1000)
	updateNews()
}