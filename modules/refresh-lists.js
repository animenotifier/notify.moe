'use strict'

let arn = require('../lib')

// Check every now and then if users have new episodes
let refreshAnimeLists = function() {
	console.log('Refreshing anime lists...')

	let tasks = []
	arn.scan('Users', function(user) {
		if(!arn.isActiveUser(user))
			return

		tasks.push(arn.getAnimeListAsync(user).then(json => {
			console.log(user.nick, json)
		}).catch(error => {
			console.error(user.nick, error)
		}))
	}, function() {
		Promise.all(tasks).then(() => {
			console.log('Finished updating all anime lists')
		})
	})
}

module.exports = function(aero, callback) {
	arn.animeListCacheTime = 300 * 1000
	setInterval(refreshAnimeLists, arn.animeListCacheTime)
}