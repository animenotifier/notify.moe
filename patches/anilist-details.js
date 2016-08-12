let aero = require('aero')
let arn = require('../lib')
let request = require('request-promise')
let Promise = require('bluebird')
let database = require('../modules/database')
let RateLimiter = require('limiter').RateLimiter
let limiter = new RateLimiter(1, 500)

database(aero, Promise.coroutine(function*(error) {
	yield arn.listProviders.AniList.authorize()

	arn.forEach('Anime', anime => {
		// Skip anime that have been imported already
		if(anime.description || anime.totalEpisodes || anime.duration)
			return

		limiter.removeTokens(1, function() {
			arn.listProviders.AniList.getAnimeDetails(anime.id).then(details => {
				arn.set('Anime', anime.id, details).then(() => console.log(`Finished importing anime ${anime.id}`))
			})
		})
	})
}))