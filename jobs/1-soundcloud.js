'use strict'

let SC = require('node-soundcloud')
let Promise = require('bluebird')
let natural = require('natural')

Promise.promisifyAll(SC)

SC.init({
	id: arn.apiKeys.soundcloud.clientID,
	secret: arn.apiKeys.soundcloud.clientSecret,
	uri: 'https://notify.moe/soundcloud/callback'
})

let findTracksForAnime = anime => {
	let searchTermOpening = anime.title.romaji + ' Opening'
	let opening = null
	let tmp = null

	return SC.getAsync('/tracks', {
		q: searchTermOpening
	}).then(tracks => {
		if(!tracks || !Array.isArray(tracks) || tracks.length === 0)
			return

		tmp = tracks

		tracks = tracks.map(track => {
			track.similarity = natural.JaroWinklerDistance(searchTermOpening, track.title)
			return track
		})

		const similarityThreshold = 0.82

		tracks.sort((a, b) => {
			if(a.similarity >= similarityThreshold && b.similarity >= similarityThreshold)
				return a.likes_count > b.likes_count ? -1 : 1

			return a.similarity > b.similarity ? -1 : 1
		})

		opening = tracks[0]

		if(opening.similarity >= similarityThreshold && opening.likes_count >= 2 && opening.title.toLowerCase().indexOf('opening') !== -1 && opening.title.toLowerCase().indexOf(anime.title.romaji.toLowerCase()) !== -1) {
			let tracks = {
				opening
			}

			return arn.set('Anime', anime.id, {
				tracks
			}).then(() => {
				return arn.updateAnimePage(anime)
			}).then(() => {
				console.log(`Saved opening track for "${anime.title.romaji}" (${opening.title})`)
			})
		}
	}).catch(error => {
		console.error(error, error.stack)
		console.log(tmp)
	})
}

let updateAnimeTracks = () => {
	console.log('Updating anime tracks...')

	arn.forEach('Anime', anime => {
		if(anime.id !== 21256 && anime.id !== 21234)
			return

		arn.networkLimiter.removeTokens(1, () => {
			findTracksForAnime(anime)
		})
	})
}

arn.repeatedly(24 * 60 * 60, () => {
	updateAnimeTracks()
})