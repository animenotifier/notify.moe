'use strict'

let arn = require('../lib')
let RateLimiter = require('limiter').RateLimiter
let limiter = new RateLimiter(1, 100)

// Check every now and then if users have new episodes
let refreshAnimeLists = function() {
	console.log('Refreshing anime lists...')

	arn.scan('Users', function(user) {
		if(!arn.isActiveUser(user))
			return

		limiter.removeTokens(1, function() {
			arn.getAnimeListAsync(user).then(animeList => {
				// ...
			}).catch(error => {
				console.error(`Error when automatically updating the anime list of ${user.nick}:`, error)
			})
		})
	}, function() {
		// ...
	})
}

module.exports = function(aero, callback) {
	arn.animeListCacheTime = 5 * 60 * 1000
	setInterval(refreshAnimeLists, arn.animeListCacheTime)
}