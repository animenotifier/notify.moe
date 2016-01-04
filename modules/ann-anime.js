'use strict'

let request = require('request-promise')
let xml2js = require('xml2js')
let Promise = require('bluebird')
let arn = require('../lib')

module.exports = function(aero) {
	/*let xmlParser = new xml2js.Parser({
		explicitArray: false,	// Don't put single nodes into an array
		ignoreAttrs: false,
		trim: true,
		normalize: true,
		explicitRoot: false
	})

	Promise.promisifyAll(xmlParser)

	let getInfo = function(ann, infoName, textNode) {
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

					if(ann.episode) {
						ann.episode.forEach(episode => {
							episodes[episode.$.num] = {
								title: episode.title._,
								language: episode.title.$.lang.toLowerCase()
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
						image: picture ? picture.img.reduce((a, b) => a.$.width > b.$.width ? a : b).$ : '',
						episodes,
						websites: ann.info.filter(info => info.$.type === 'Official website').map(website => {
							return {
								title: website._,
								url: website.$.href,
								type: website.$.type,
								lang: website.$.lang
							}
						}),
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
				let animeIds = json.item.map(item => item.id).join('/')

				if(!animeIds)
					return

				return animeIds
			})
		}).catch(error => {
			console.log('Error getting list of anime:', error)
		})
	}

	getAnimeIds('http://www.animenewsnetwork.com/encyclopedia/reports.xml?id=155&type=anime&nlist=50')
	.then(animeIds => importAnime(`http://cdn.animenewsnetwork.com/encyclopedia/api.xml?title=${animeIds}`))
	.then(animeList => {
		//console.log(require('util').inspect(animeList, false, null))

		animeList.forEach(anime => {
			if(!anime.id || !anime.title)
				return

			arn.setAsync('Anime', anime.id, anime).then(() => console.log('Imported anime: ' + anime.title))
		})
	}).catch(error => {
		console.log('ANN import error:', error)
	})
	//http://www.animenewsnetwork.com/encyclopedia/reports.xml?id=155&type=anime&nlist=all*/
}