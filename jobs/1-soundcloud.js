'use strict'

let SC = require('node-soundcloud')
let Promise = require('bluebird')
let natural = require('natural')
let chalk = require('chalk')

Promise.promisifyAll(SC)

SC.init({
	id: arn.apiKeys.soundcloud.clientID,
	secret: arn.apiKeys.soundcloud.clientSecret,
	uri: 'https://notify.moe/soundcloud/callback'
})

let findTracksForAnime = anime => {
	if(anime.tracks && anime.tracks.opening)
		return

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

			// Assign tracks to this object so that the anime page already has the info
			if(!anime.tracks)
				anime.tracks = tracks
			else
				anime.tracks = Object.assign({}, anime.tracks, tracks)

			return arn.set('Anime', anime.id, {
				tracks
			}).then(() => {
				// We add a delay to let the database catch up
				return Promise.delay(100).then(() => arn.updateAnimePage(anime))
			}).then(() => {
				console.log(chalk.green('✔'), `Saved opening track for ${chalk.cyan(anime.title.romaji)} (${opening.title})`)
			})
		}
	}).catch(error => {
		console.error(error, error.stack)
		console.log(tmp)
	})
}

let updateAnimeTracks = () => {
	console.log(chalk.yellow('✖'), 'Updating anime tracks...')

	arn.forEach('Anime', anime => {
		arn.networkLimiter.removeTokens(1, () => {
			findTracksForAnime(anime)
		})
	})
}

arn.repeatedly(24 * 60 * 60, () => {
	updateAnimeTracks()
})