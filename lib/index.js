'use strict'

let aero = require('aero')
let Promise = require('bluebird')
let EventEmitter = require('events').EventEmitter
let apiKeys = require('../security/api-keys.json')
let request = require('request-promise')

let arn = {
	events: new EventEmitter(),
	push: require('./push')
}

arn.listProviders = {
	AniList: require('./providers/animelist/AniList'),
	AnimePlanet: require('./providers/animelist/AnimePlanet'),
	HummingBird: require('./providers/animelist/HummingBird'),
	MyAnimeList: require('./providers/animelist/MyAnimeList')
}

arn.animeProviders = {
	Nyaa: require('./providers/anime/Nyaa')
}

arn.airingDateProviders = {
	AniList: arn.listProviders.AniList
}

arn.get = function(set, key, callback) {
	aero.db.get({
		ns: 'arn',
		set: set,
		key: key
	}, function(error, record, metadata, key) {
		if(callback)
			callback(error.code !== 0 ? error : undefined, record)
		else if(error.code !== 0)
			console.error(error)
	})
}

arn.set = function(set, key, obj, callback) {
	let aerospikeKey = {
		ns: 'arn',
		set: set,
		key: key
	}

	aero.db.put(aerospikeKey, obj, function(error) {
		if(callback)
			callback(error.code !== 0 ? error : undefined, obj)
		else if(error.code !== 0)
			console.error(error)
	})
}

arn.remove = function(set, key, callback) {
	aero.db.remove({
		ns: 'arn',
		set: set,
		key: key
	}, function(error, key) {
		if(callback)
			callback(error.code !== 0 ? error : undefined, key)
		else if(error.code !== 0)
			console.error(error)
	})
}

arn.registerNewUser = function(user, customTask) {
	return Promise.all([
		arn.setAsync('NickToUser', user.nick, { userId: user.id }),
		arn.setAsync('EmailToUser', user.email, { userId: user.id }),
		/*arn.setAsync('AnimeList', user.id, {
			watching: [],
			completed: [],
			onHold: [],
			dropped: [],
			planToWatch: []
		}),*/
		customTask
	]).then(function() {
		arn.events.emit('new user', user)
	})
}

arn.getUser = function(userId, callback) {
	aero.db.get({
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
	aero.db.get({
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
	aero.db.put(
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
	let scan = aero.db.query('arn', set, {
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

	arn.getAsync('NickToUser', newNick)
	.then(record => {
		callback('Username is already taken.')
	})
	.catch(error => {
		user.nick = newNick

		Promise.all([
			arn.removeAsync('NickToUser', oldNick),
			arn.setAsync('NickToUser', newNick, { userId: user.id }),
			arn.setUserAsync(user.id, user)
		]).then(() => {
			callback()
		}).catch(error => {
			callback(error)
		})
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

module.exports = arn