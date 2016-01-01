'use strict'

let arn = {}
let aero = require('aero')
let Promise = require('bluebird')

arn.listProviders = {
	AniList: require('./providers/animelist/AniList'),
	AnimePlanet: require('./providers/animelist/AnimePlanet'),
	HummingBird: require('./providers/animelist/HummingBird'),
	MyAnimeList: require('./providers/animelist/MyAnimeList')
}

arn.animeProviders = {
	// ...
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
			callback(error.code !== 0 ? error : undefined)
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
	])
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
				callback(error.code !== 0 ? error : undefined)
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

Promise.promisifyAll(arn)
module.exports = arn