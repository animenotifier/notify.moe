'use strict'

let aero = require('aero')
let Promise = require('bluebird')
let EventEmitter = require('events').EventEmitter
let apiKeys = require('../security/api-keys.json')
let request = require('request-promise')
let sortAlgorithms = require('./sort-algorithms')
let natural = require('natural')

let arn = {
	events: new EventEmitter(),
	listProviders: {
		AniList: require('./providers/list/AniList'),
		AnimePlanet: require('./providers/list/AnimePlanet'),
		HummingBird: require('./providers/list/HummingBird'),
		MyAnimeList: require('./providers/list/MyAnimeList')
	},
	animeProviders: {
		Nyaa: require('./providers/download/Nyaa')
		//AnimeTwist: require('./providers/watch/AnimeTwist')
	},
	cacheAnimeLists: true
}

arn.airingDateProviders = {
	AniList: arn.listProviders.AniList
}

require('./database').init(arn)
require('./notifications').init(arn)

arn.registerNewUser = function(user, customTask) {
	return Promise.all([
		arn.set('NickToUser', user.nick, { userId: user.id }),
		arn.set('EmailToUser', user.email, { userId: user.id }),
		customTask
	]).then(function() {
		arn.events.emit('new user', user)
	})
}

arn.getUserByNick = Promise.coroutine(function*(nick) {
	let record = yield arn.get('NickToUser', nick)
	return arn.get('Users', record.userId)
})

arn.scan = Promise.promisify(function(set, func, callback) {
	let scan = arn.db.query('arn', set, {
	    concurrent: true,
	    nobins: false
	})

	let stream = scan.execute()

	stream.on('data', function(record) {
		func(record.bins)
	})

	stream.on('error', function(error) {
		console.error('Error occured while scanning:', error, error.stack)
	})

	if(callback)
		stream.on('end', callback)
})

arn.on = function(eventName, func) {
	arn.events.on(eventName, func)
}

arn.changeNick = function(user, newNick) {
	let oldNick = user.nick

	if(oldNick === newNick)
		return Promise.resolve()

	return arn.get('NickToUser', newNick).then(record => {
		return Promise.reject('Username is already taken.')
	}).catch(error => {
		user.nick = newNick

		return Promise.all([
			arn.remove('NickToUser', oldNick),
			arn.set('NickToUser', newNick, { userId: user.id }),
			arn.set('Users', user.id, user)
		])
	})
}

arn.refreshAnimeList = Promise.promisify(function(user, listProvider, animeProvider, airingDateProvider, listProviderSettings, cacheKey, callback) {
	listProvider.getAnimeList(listProviderSettings.userName, (error, watching) => {
		if(error) {
			callback(error, watching)
			return
		}

		let asyncTasks = []

		watching.forEach(entry => {
			entry.animeProvider = {
				url: null,
				nextEpisodeUrl: null,
				videoUrl: null
			}

			if(listProvider === airingDateProvider && airingDateProvider.getAiringDateById)
				asyncTasks.push(airingDateProvider.getAiringDateById(entry, entry.providerId))
			else
				asyncTasks.push(airingDateProvider.getAiringDate(entry))

			if(animeProvider)
				asyncTasks.push(animeProvider.getAnimeInfo(entry))
		})

		Promise.all(asyncTasks).then(() => {
			watching.sort(sortAlgorithms[user.sortBy])

			let animeList = {
				user: user.nick,
				userId: user.id,
				listProvider: user.providers.list,
				listUrl: listProvider.getAnimeListUrl(listProviderSettings.userName),
				watching,
				cacheKey,
				generated: (new Date()).toISOString()
			}

			// Cache it
			arn.set('AnimeLists', user.id, animeList).then(() => {
				callback(undefined, animeList)
			})
		}).catch(error => {
			callback(error, null)
		})
	})
})

arn.getAnimeList = function(user, clearCache, callback) {
	let listProviderName = user.providers.list
	let listProvider = arn.listProviders[listProviderName]
	let animeProviderName = user.providers.anime
	let animeProvider = arn.animeProviders[animeProviderName]
	let airingDateProvider = arn.airingDateProviders[user.providers.airingDate]
	let listProviderSettings = user.listProviders[listProviderName]

	if(!listProvider)
		callback(new Error('Invalid list provider'))

	if(!listProviderSettings || !listProviderSettings.userName)
		callback(new Error(`${listProviderName} username has not been specified`))

	let cacheKey = listProviderName + ':' + listProviderSettings.userName + ':' + user.sortBy

	let refresh = (oldAnimeList) => {
		return arn.refreshAnimeList(user, listProvider, animeProvider, airingDateProvider, listProviderSettings, cacheKey).then(animeList => {
			if(!oldAnimeList)
				return animeList

			// Compare to check if we can send notifications
			if(Object.keys(user.devices).length > 0) {
				animeList.watching.forEach(anime => {
					let oldAnime = oldAnimeList.watching.find(e => e.providerId === anime.providerId)

					if(!oldAnime)
						return

					// Send push notification
					if(anime.episodes && oldAnime.episodes && anime.episodes.available > oldAnime.episodes.available && anime.episodes.available === anime.episodes.next) {
						arn.sendNotification(user, {
							title: anime.title,
							icon: anime.image,
							message: `Episode ${anime.episodes.available} was just released`,
							tag: 'new-episode'
						})
					}
				})
			}

			return animeList
		})
		.then(animeList => callback(undefined, animeList))
		.catch(callback)
	}

	arn.get('AnimeLists', user.id).then(animeList => {
		let now = new Date()
		let generated = new Date(animeList.generated)

		if(arn.cacheAnimeLists && !clearCache && cacheKey === animeList.cacheKey && now.getTime() - generated.getTime() < arn.animeListCacheTime) {
			callback(undefined, animeList)
		} else {
			refresh(animeList)
		}
	}).catch(error => {
		if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND') {
			refresh()
		} else {
			callback(error)
		}
	})
}

// -----------------------------------------------------
Promise.promisifyAll(arn)
// -----------------------------------------------------

let autoCorrectUserNames = [
	/anilist.co\/user\/(.*)/,
	/anilist.co\/animelist\/(.*)/,
	/hummingbird.me\/users\/(.*?)\/library/,
	/hummingbird.me\/users\/(.*)/,
	/anime-planet.com\/users\/(.*?)\/anime/,
	/anime-planet.com\/users\/(.*)/,
	/myanimelist.net\/profile\/(.*)/,
	/myanimelist.net\/animelist\/(.*?)\?/,
	/myanimelist.net\/animelist\/(.*)/
]

arn.fixListProviderUserName = function(userName) {
	userName = userName.trim()

	for(let regex of autoCorrectUserNames) {
		let match = regex.exec(userName)

		if(match !== null) {
			userName = match[1]
			break
		}
	}

	return userName
}

arn.fixNick = function(nick) {
	nick = nick.replace(/[\W\s\d]/g, '')

	if(nick)
		nick = nick[0].toUpperCase() + nick.substring(1)

	return nick
}

arn.getLocation = function(user) {
	let locationAPI = `http://api.ipinfodb.com/v3/ip-city/?key=${apiKeys.ipInfoDB.clientID}&ip=${user.ip}&format=json`
	return request(locationAPI).then(JSON.parse)
}

arn.isActiveUser = function(user) {
	if(user.nick.startsWith('g'))
		return false

	if(user.nick.startsWith('fb'))
		return false

	let listProviderName = user.providers.list

	if(!listProviderName)
		return false

	let listProvider = user.listProviders[listProviderName]

	if(!listProvider || !listProvider.userName)
		return false

	return true
}

arn.getAnimeListByNickAsync = function(nick, clearCache) {
	return arn.getUserByNick(nick).then(user => arn.getAnimeListAsync(user, clearCache)).catch(error => {
		console.error(error.stack)

		if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND')
			return Promise.reject(`User '${nick}' not found`)

		if(error.message)
			return Promise.reject(error.message)

		return Promise.reject(error.toString())
	})
}

arn.getAnimeIdBySimilarTitle = Promise.promisify(function(anime, listProviderName, callback) {
	if(!anime || !anime.providerId)
		return callback(undefined, null)

	let bucket = 'Match' + listProviderName

	// Look up cached or corrected version by ID
	arn.get(bucket, anime.providerId).then(match => {
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
		arn.set(bucket, anime.providerId, bestResult)

		return callback(undefined, bestResult)
	})
})

module.exports = arn