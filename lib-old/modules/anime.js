let Promise = require('bluebird')
let natural = require('natural')

arn.getAnimeIdBySimilarTitle = Promise.promisify(function(anime, listProviderName, callback) {
	if(!anime || !anime.providerId)
		return callback(undefined, null)

	let bucket = 'Match' + listProviderName

	// Look up cached or corrected version by ID
	arn.db.get(bucket, anime.providerId).then(match => {
		callback(undefined, match)
	}).catch(error => {
		let search = anime.title

		if(!search || search === null)
			return callback(undefined, null)

		if(!arn.animeToId) {
			console.error('Anime to ID index has not been built yet.')
			return callback(undefined, null)
		}

		let arnTitles = Object.keys(arn.animeToId)
		let searchResults = arnTitles.map(title => {
			return {
				id: arn.animeToId[title],
				providerId: anime.providerId,
				title,
				providerTitle: search,
				similarity: natural.JaroWinklerDistance(search, title)
			}
		})

		searchResults.sort((a, b) => {
			return a.similarity < b.similarity ? 1 : -1
		})

		let bestResult = searchResults[0]

		// Save in database
		arn.db.set(bucket, anime.providerId, bestResult)

		return callback(undefined, bestResult)
	})
})