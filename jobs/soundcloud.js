let SC = require('node-soundcloud')
let natural = require('natural')

Promise.promisifyAll(SC)

SC.init({
	id: arn.apiKeys.soundcloud.clientID,
	secret: arn.apiKeys.soundcloud.clientSecret,
	uri: 'https://notify.moe/soundcloud/callback'
})

let findTracksForAnime = anime => {
	//if(anime.tracks && anime.tracks.opening)
	//	return

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
				opening: {
					uri: opening.uri,
					title: opening.title,
					similarity: opening.similarity,
					likes: opening.likes_count,
					plays: opening.playback_count,
					permalink: opening.permalink_url
				}
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
		} else {
			console.log(chalk.red('✖'), `No tracks found for ${chalk.cyan(anime.title.romaji)}`)
		}
	}).catch(error => {
		console.error(error, error.stack)
		console.log(tmp)
	})
}

let updateAnimeTracks = coroutine(function*() {
	console.log(chalk.yellow('✖'), 'Updating anime tracks...')

	for(let anime of arn.animeList) {
		yield Promise.delay(1500)
		findTracksForAnime(anime)
	}
})

arn.repeatedly(24 * hours, updateAnimeTracks)