'use strict'

let Promise = require('bluebird')

exports.init = (arn) => {
	arn.get = Promise.promisify(function(set, key, callback) {
		arn.db.get({
			ns: 'arn',
			set: set,
			key: key
		}, function(error, record, metadata, key) {
			if(callback)
				callback(error.code !== 0 ? error : undefined, record)
			else if(error.code !== 0)
				console.error(error)
		})
	})

	arn.set = Promise.promisify(function(set, key, obj, callback) {
		let aerospikeKey = {
			ns: 'arn',
			set: set,
			key: key
		}

		arn.db.put(aerospikeKey, obj, function(error) {
			if(callback)
				callback(error.code !== 0 ? error : undefined, obj)
			else if(error.code !== 0)
				console.error(error)
		})
	})

	arn.remove = Promise.promisify(function(set, key, callback) {
		arn.db.remove({
			ns: 'arn',
			set: set,
			key: key
		}, function(error, key) {
			if(callback)
				callback(error.code !== 0 ? error : undefined, key)
			else if(error.code !== 0)
				console.error(error)
		})
	})
}