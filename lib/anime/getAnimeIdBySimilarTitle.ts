import * as arn from 'lib'
import * as natural from 'natural'

export async function getAnimeIdBySimilarTitle(anime, listProviderName: string) {
	if(!anime || !anime.providerId)
		return null

	let bucket = 'Match' + listProviderName

	// Look up cached or corrected version by ID
	return await arn.db.get(bucket, anime.providerId).catch(error => {
		let search = anime.title

		if(!search || search === null)
			return null

		if(!arn.titleToId) {
			console.error('Anime to ID index has not been built yet.')
			return null
		}

		let arnTitles = Object.keys(arn.titleToId)
		let searchResults = arnTitles.map(title => {
			return {
				id: arn.getIdByTitle(title),
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

		return bestResult
	})
}