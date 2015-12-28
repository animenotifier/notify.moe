'use strict'

let arn = {}
let aero = require('aero')
let Promise = require('bluebird')

arn.get = function(set, key, callback) {
	aero.db.get({
		ns: 'arn',
		set: set,
		key: key
	}, function(error, record, metadata, key) {
		callback(error.code !== 0 ? error : undefined, record)
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

arn.getUser = function(userId, callback) {
	aero.db.get({
		ns: 'arn',
		set: 'Users',
		key: userId
	}, function(error, user, metadata, key) {
		callback(error.code !== 0 ? error : undefined, user)
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
		callback ? callback : function(error) {
			if(error.code !== 0)
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