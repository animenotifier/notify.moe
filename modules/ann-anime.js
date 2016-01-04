'use strict'

let request = require('request-promise')
let xml2js = require('xml2js')
let Promise = require('bluebird')
let arn = require('../lib')

module.exports = function(aero) {
	let xmlParser = new xml2js.Parser({
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
		request(url)
		.then(xml => {
			return xmlParser.parseStringAsync(xml).then(json => {
				let result = json.anime.map(ann => {
					let episodes = []

					ann.episode.forEach(episode => {
						episodes[episode.$.num] = {
							title: episode.title._,
							language: episode.title.$.lang.toLowerCase()
						}
					})

					let anime = {
						id: ann.$.id,
						name: ann.$.name,
						altNames: [],
						type: ann.$.type,
						precision: ann.$.precision,
						summary: getInfo(ann, 'Plot Summary', true),
						image: getInfo(ann, 'Picture').img.reduce((a, b) => a.$.width > b.$.width ? a : b).$,
						episodes,
						websites: ann.info.filter(info => info.$.type === 'Official website').map(website => {
							return {
								title: website._,
								url: website.$.href,
								type: website.$.type,
								lang: website.$.lang
							}
						}),
						staff: [],
						cast: [],
						credit: []
					}

					return anime
				})

				console.log(require('util').inspect(result, false, null))
			})
		})
		.catch(error => {
			console.log('Error importing anime:', error)
		})
	}

	/*let importAnimeList = url => {
		request(url)
		.then(xml => {
			return xmlParser.parseStringAsync(xml).then(json => {
				let animeIds = json.item.map(item => item.id).join('/')

				if(!animeIds)
					return

				importAnime(`http://cdn.animenewsnetwork.com/encyclopedia/api.xml?title=${animeIds}`)
			})
		})
		.catch(error => {
			console.log('Error importing anime:', error)
		})
	}

	importAnimeList('http://www.animenewsnetwork.com/encyclopedia/reports.xml?id=155&type=anime&nlist=2')*/

	importAnime(`http://cdn.animenewsnetwork.com/encyclopedia/api.xml?title=15118/16134`)
	//http://www.animenewsnetwork.com/encyclopedia/reports.xml?id=155&type=anime&nlist=all
}