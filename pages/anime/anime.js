'use strict'

let Promise = require('bluebird')
let gravatar = require('gravatar')
let striptags = require('striptags')
let popularAnimeCached = []

let updatePopularAnime = function() {
	let popularAnime = []

	return arn.forEach('Anime', anime => {
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
			providers[listProviderName] = []

			return match => {
				if(match.id === animeId)
					providers[listProviderName].push(match)
			}
		}

		let userQueryTasks = []

		let otherTasks = {
			HummingBird: arn.forEach('MatchHummingBird', createScanFunction('HummingBird')),
			MyAnimeList: arn.forEach('MatchMyAnimeList', createScanFunction('MyAnimeList')),
			AnimePlanet: arn.forEach('MatchAnimePlanet', createScanFunction('AnimePlanet')),
			Nyaa: arn.get('AnimeToNyaa', animeId).catch(error => undefined),
			Watching: arn.forEach('AnimeLists', list => {
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

		let sortMatches = (a, b) => {
			if((a.edited && b.edited) || (!a.edited && !b.edited))
				return a.similarity < b.similarity ? 1 : -1

			if(a.edited)
				return -1

			// b edited
			return 1
		}

		arn.get('Anime', animeId).then(anime => {
			anime.description = striptags(anime.description, ['br'])

			Promise.props(otherTasks).then(result => {
				providers.MyAnimeList.sort(sortMatches)
				providers.HummingBird.sort(sortMatches)
				providers.AnimePlanet.sort(sortMatches)

				providers.MyAnimeList = providers.MyAnimeList[0]
				providers.HummingBird = providers.HummingBird[0]
				providers.AnimePlanet = providers.AnimePlanet[0]

				providers.Nyaa = result.Nyaa

				Promise.all(userQueryTasks).then(usersWatching => {
					response.render({
						user,
						anime,
						providers,
						usersWatching,
						nyaa: arn.animeProviders.Nyaa,
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