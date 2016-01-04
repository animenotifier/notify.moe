'use strict'

let fs = require('fs')
let request = require('request-promise')
let xml2js = require('xml2js')
let Promise = require('bluebird')
let arn = require('../lib')
let RateLimiter = require('limiter').RateLimiter
let limiter = new RateLimiter(1, 1100)
let database = require('../modules/database')
let aero = require('aero')

let xmlParser = new xml2js.Parser({
	explicitArray: true,	// Don't put single nodes into an array
	ignoreAttrs: false,
	trim: true,
	normalize: true,
	explicitRoot: false
})

Promise.promisifyAll(xmlParser)

let getInfo = function(ann, infoName, textNode) {
	if(!Array.isArray(ann.info))
		return ''

	let values = ann.info.filter(info => info.$.type === infoName)

	if(textNode)
		values = values.map(node => node._)

	return values ? values[0] : ''
}

// Anime database
let importAnime = url => {
	return request(url).then(xml => {
		return xmlParser.parseStringAsync(xml).then(json => {
			let result = json.anime.map(ann => {
				let episodes = []

				if(Array.isArray(ann.episode)) {
					ann.episode.forEach(episode => {
						episodes[episode.$.num] = {
							title: episode.title._,
							language: episode.title.$ ? episode.title.$.lang.toLowerCase() : ''
						}
					})
				}

				let picture = getInfo(ann, 'Picture')
				let summary = getInfo(ann, 'Plot Summary', true)

				let anime = {
					id: ann.$.id ? ann.$.id : -1,
					title: ann.$.name ? ann.$.name : '',
					altNames: [],
					type: ann.$.type ? ann.$.type : '',
					precision: ann.$.precision,
					summary: summary ? summary : '',
					image: (picture && picture.img && picture.img.length > 0) ? picture.img.reduce((a, b) => a.$.width > b.$.width ? a : b).$ : '',
					episodes,
					websites: Array.isArray(ann.info) ? ann.info.filter(info => info.$.type === 'Official website').map(website => {
						return {
							title: website._,
							url: website.$.href,
							type: website.$.type,
							lang: website.$.lang
						}
					}) : [],
					genres: [],
					staff: [],
					cast: [],
					credit: []
				}

				return anime
			})

			return result
		})
	}).catch(error => {
		console.log('Error importing anime:', error)
	})
}

let getAnimeIds = url => {
	return request(url).then(xml => {
		return xmlParser.parseStringAsync(xml).then(json => {
			return json.item.map(item => item.id)
		})
	}).catch(error => {
		console.log('Error getting list of anime:', error)
	})
}

database(aero, function(error) {
	getAnimeIds('http://www.animenewsnetwork.com/encyclopedia/reports.xml?id=155&type=anime&nlist=all')
	.then(animeIds => {
		console.log('Found ' + animeIds.length + ' anime')
		return animeIds
	})
	.then(animeIds => {
		for(let i = 0; i < animeIds.length; i += 50) {
			let h = i
			let part = animeIds.slice(h, h + 50)
			let partString = part.join('/')

			limiter.removeTokens(1, function() {
				console.log('Requesting anime from', h, 'to', h + 49)

				importAnime(`http://cdn.animenewsnetwork.com/encyclopedia/api.xml?title=${partString}`)
				.then(animeList => {
					animeList.forEach(anime => {
						if(!anime.id || !anime.title)
							return

						arn.setAsync('Anime', anime.id, anime).catch(error => {
							console.error('Error saving anime', anime.id)
						}) //.then(() => console.log('Imported anime: ' + anime.title))
					})
				}).catch(error => {
					console.log('ANN import error:', error)
				})
			});
		}
	})
	.then(() => console.log('Finished import'))
	.catch(error => {
		console.log('ANN import error:', error)
	})
})