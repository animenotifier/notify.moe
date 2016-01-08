'use strict'

let aero = require('aero')
let Promise = require('bluebird')
let EventEmitter = require('events').EventEmitter
let apiKeys = require('../security/api-keys.json')
let request = require('request-promise')
let sortAlgorithms = require('./sort-algorithms')

let arn = {
	events: new EventEmitter(),
	listProviders: {
		AniList: require('./providers/list/AniList'),
		AnimePlanet: require('./providers/list/AnimePlanet'),
		HummingBird: require('./providers/list/HummingBird'),
		MyAnimeList: require('./providers/list/MyAnimeList')
	},
	animeProviders: {
		Nyaa: require('./providers/anime/Nyaa')
	}
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

arn.getUser = function(userId, callback) {
	arn.db.get({
		ns: 'arn',
		set: 'Users',
		key: userId
	}, function(error, user, metadata, key) {
		if(callback)
			callback(error.code !== 0 ? error : undefined, user)
		else if(error.code !== 0)
			console.error(error)
	})
}

arn.getUserByNick = function(nick, callback) {
	arn.db.get({
		ns: 'arn',
		set: 'NickToUser',
		key: nick
	}, function(error, record, metadata, key) {
		if(error.code === 0)
			arn.getUser(record.userId, callback)
		else
			callback(error, undefined)
	})
}

arn.setUser = function(userId, user, callback) {
	arn.db.put(
		{
			ns: 'arn',
			set: 'Users',
			key: userId
		},
		user,
		function(error) {
			if(callback)
				callback(error.code !== 0 ? error : undefined, user)
			else if(error.code !== 0)
				console.error(error)
		}
	)
}

arn.scan = function(set, func, callback) {
	let scan = arn.db.query('arn', set, {
	    concurrent: true,
	    nobins: false
	})

	let stream = scan.execute()

	stream.on('data', function(record) {
		func(record.bins)
	})

	stream.on('error', function(error) {
		console.error('Error occured while scanning:', error)
	})

	if(callback)
		stream.on('end', callback)
}

arn.on = function(eventName, func) {
	arn.events.on(eventName, func)
}

arn.changeNick = function(user, newNick, callback) {
	let oldNick = user.nick

	if(oldNick === newNick)
		return callback()

	arn.get('NickToUser', newNick)
	.then(record => {
		callback('Username is already taken.')
	})
	.catch(error => {
		user.nick = newNick

		Promise.all([
			arn.remove('NickToUser', oldNick),
			arn.set('NickToUser', newNick, { userId: user.id }),
			arn.setUserAsync(user.id, user)
		]).then(() => {
			callback()
		}).catch(error => {
			callback(error)
		})
	})
}

arn.getAnimeList = function(user, callback) {
	let listProviderName = user.providers.list
	let listProvider = arn.listProviders[listProviderName]
	let animeProviderName = user.providers.anime
	let animeProvider = arn.animeProviders[animeProviderName]
	let airingDateProvider = arn.airingDateProviders[user.providers.airingDate]
	let listProviderSettings = user.listProviders[listProviderName]

	if(!listProvider)
		callback('Invalid list provider')

	if(!listProviderSettings || !listProviderSettings.userName)
		callback(`${listProviderName} username has not been specified`)

	let cacheKey = listProviderName + ':' + listProviderSettings.userName + ':' + user.sortBy

	let refreshList = () => {
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
					listProvider: listProviderName,
					listUrl: listProvider.getAnimeListUrl(listProviderSettings.userName),
					watching,
					cacheKey,
					generated: (new Date()).toISOString()
				}

				callback(undefined, animeList)

				// Cache it
				arn.set('AnimeLists', user.id, animeList)
			}).catch(error => {
				callback(error, null)
			})
		})
	}

	arn.get('AnimeLists', user.id).then(animeList => {
		let now = new Date()
		let generated = new Date(animeList.generated)

		if(cacheKey === animeList.cacheKey && now.getTime() - generated.getTime() < arn.animeListCacheTime)
			callback(undefined, animeList)
		else
			refreshList()
	}).catch(error => {
		if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND')
			refreshList()
		else
			callback(error)
	})
}

// -----------------------------------------------------
Promise.promisifyAll(arn)
// -----------------------------------------------------

arn.getLocation = function(user) {
	let locationAPI = `http://api.ipinfodb.com/v3/ip-city/?key=${apiKeys.ipInfoDB.clientID}&ip=${user.ip}&format=json`
	return request(locationAPI).then(JSON.parse)
}

arn.fixNick = function(nick) {
	nick = nick.replace(/[\W\s\d]/g, '')

	if(nick)
		nick = nick[0].toUpperCase() + nick.substring(1)

	return nick
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

arn.getAnimeListByNickAsync = function(nick) {
	return arn.getUserByNickAsync(nick).then(arn.getAnimeListAsync).catch(error => {
		if(error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND')
			return Promise.reject(`User '${nick}' not found`)

		if(error.message)
			return Promise.reject(error.message)

		return Promise.reject(error.toString())
	})
}

module.exports = arn