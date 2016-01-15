'use strict'

let Promise = require('bluebird')
let gravatar = require('gravatar')
let popularAnimeCached = []

let updatePopularAnime = function() {
	let popularAnime = []

	return arn.scan('Anime', anime => {
		if(anime.watching)
			popularAnime.push(anime)
	}).then(() => {
		popularAnime.sort((a, b) => a.watching < b.watching ? 1 : -1)

		if(popularAnime.length > maxPopularAnime)
			popularAnime.length = maxPopularAnime

		return popularAnime
	})
}

setInterval(function() {
	updatePopularAnime().then(popularAnime => {
		popularAnimeCached = popularAnime
	})
}, 60 * 1000)

const maxPopularAnime = 10

exports.get = function(request, response) {
	let user = request.user
	let animeId = parseInt(request.params[0])

	if(animeId) {
		let providers = {}

		let createScanFunction = function(listProviderName) {
			return match => {
				if(match.id === animeId)
					providers[listProviderName] = match
			}
		}

		let userQueryTasks = []

		let otherTasks = {
			HummingBird: arn.scan('MatchHummingBird', createScanFunction('HummingBird')),
			MyAnimeList: arn.scan('MatchMyAnimeList', createScanFunction('MyAnimeList')),
			AnimePlanet: arn.scan('MatchAnimePlanet', createScanFunction('AnimePlanet')),
			Watching: arn.scan('AnimeLists', list => {
				if(!list.userId)
					return

				if(list.watching.find(anime => anime.id === animeId)) {
					userQueryTasks.push(arn.get('Users', list.userId).then(user => {
						user.gravatarURL = gravatar.url(user.email, {s: '50', r: 'x', d: '404'}, true)
						return user
					}))
				}
			})
		}

		arn.get('Anime', animeId).then(anime => {
			Promise.props(otherTasks).then(result => {
				Promise.all(userQueryTasks).then(usersWatching => {
					response.render({
						user,
						anime,
						providers,
						usersWatching,
						canEdit: user && (user.role === 'admin' || user.role === 'editor')
					})

					// Save number of people watching
					arn.set('Anime', animeId, {
						watching: usersWatching.length
					})
				})
			})
		}).catch(error => {
			console.log(error.stack)
			response.writeHead(404)
			response.end('Anime not found.')
		})
		return
	}

	if(popularAnimeCached.length === 0) {
		updatePopularAnime().then(popularAnime => {
			popularAnimeCached = popularAnime

			response.render({
				user,
				popularAnime: popularAnimeCached,
				animeToIdCount: arn.animeToIdCount
			})
		})
	} else {
		response.render({
			user,
			popularAnime: popularAnimeCached,
			animeToIdCount: arn.animeToIdCount
		})
	}
}