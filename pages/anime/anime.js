'use strict'

let Promise = require('bluebird')
let gravatar = require('gravatar')
let striptags = require('striptags')

const maxPopularAnime = 10
const animePageCacheTime = 60 * 60 * 1000
let sourceRegEx = /\(Source: (.*?)\)/i

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
	}).then(popularAnime => {
		return arn.set('Cache', 'popularAnime', popularAnime)
	})
}

let updateAnimePage = anime => {
	console.log(`Updating anime page ${anime.id} [${anime.title.romaji}]`)

	let providers = {}

	let createScanFunction = function(listProviderName) {
		providers[listProviderName] = []

		return match => {
			if(match.id === anime.id)
				providers[listProviderName].push(match)
		}
	}

	let userQueryTasks = []

	let otherTasks = {
		HummingBird: arn.forEach('MatchHummingBird', createScanFunction('HummingBird')),
		MyAnimeList: arn.forEach('MatchMyAnimeList', createScanFunction('MyAnimeList')),
		AnimePlanet: arn.forEach('MatchAnimePlanet', createScanFunction('AnimePlanet')),
		Nyaa: arn.get('AnimeToNyaa', anime.id).catch(error => undefined),
		Watching: arn.forEach('AnimeLists', list => {
			if(!list.userId)
				return

			if(list.watching.find(entry => entry.id === anime.id)) {
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

	anime.description = striptags(anime.description)

	let summarySource = ''
	let sourceMatch = anime.description.match(sourceRegEx)

	if(sourceMatch !== null) {
		summarySource = sourceMatch[1]
		anime.description = anime.description.replace(sourceRegEx, '').trim()
	}

	return Promise.props(otherTasks).then(result => {
		providers.MyAnimeList.sort(sortMatches)
		providers.HummingBird.sort(sortMatches)
		providers.AnimePlanet.sort(sortMatches)

		providers.MyAnimeList = providers.MyAnimeList[0]
		providers.HummingBird = providers.HummingBird[0]
		providers.AnimePlanet = providers.AnimePlanet[0]

		providers.Nyaa = result.Nyaa

		return Promise.all(userQueryTasks).then(usersWatching => {
			let generated = (new Date()).toISOString()

			let animePage = {
				anime,
				providers,
				usersWatching,
				summarySource,
				generated
			}

			// Save number of people watching
			arn.set('Anime', anime.id, {
				watching: usersWatching.length,
				pageGenerated: generated
			})

			// Save anime page
			return arn.set('AnimePages', anime.id, animePage)
		})
	})
}

let updateAllAnimePages = () => {
	let now = new Date()

	return arn.forEach('Anime', anime => {
		if(anime.pageGenerated && now.getTime() - (new Date(anime.pageGenerated)).getTime() < animePageCacheTime)
			return

		arn.cacheLimiter.removeTokens(1, () => {
			updateAnimePage(anime)
		})
	})
}

// Repeatedly
//arn.repeatedly(5 * 60, updatePopularAnime)
//arn.repeatedly(5 * 60 * 60, updateAllAnimePages)

exports.get = function*(request, response) {
	let user = request.user
	let animeId = parseInt(request.params[0])

	if(animeId) {
		try {
			let animePage = yield arn.get('AnimePages', animeId)

			response.render(Object.assign({
				user,
				fixGenre: arn.fixGenre,
				nyaa: arn.animeProviders.Nyaa,
				canEdit: user && (user.role === 'admin' || user.role === 'editor')
			}, animePage))
		} catch(error) {
			console.error(error, error.stack)

			response.render({
				user,
				error: 404
			})
		}

		return
	}

	let popularAnime = yield arn.get('Cache', 'popularAnime')

	response.render({
		user,
		popularAnime,
		animeToIdCount: arn.animeToIdCount
	})
}