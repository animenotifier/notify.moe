'use strict'

let Promise = require('bluebird')
let aerospike = require('aerospike')

arn.db = aerospike.client({
    hosts: [{
        addr: '127.0.0.1',
        port: 3000
    }],
    log: {
        level: aerospike.log.INFO
    },
    policies: {
        timeout: 5000
    }
})

arn.connectDB = Promise.promisify(callback => {
	arn.db.connect(response => {
	    if(response.code === 0) {
	        console.log('Successfully connected to database!')

			arn.events.emit('database ready')

			if(callback)
				callback(undefined)
	    } else {
	        console.error('Couldn\'t connect to database!')
	        console.error(response)

			if(callback)
				callback(response)
	    }
	})
})

arn.db.ready = arn.connectDB()

arn.get = Promise.promisify(function(set, key, callback) {
	arn.db.get({
		ns: 'arn',
		set: set,
		key: key
	}, function(error, record, metadata, key) {
		if(callback)
			callback(error.code !== 0 ? error : undefined, record)
		else if(error.code !== 0)
			console.error(error, error.stack)
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
			console.error(error, error.stack)
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
			console.error(error, error.stack)
	})
})

arn.forEach = Promise.promisify(function(set, func, callback) {
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

arn.filter = (set, func) => {
	let records = []

	return arn.forEach(set, record => {
		if(func(record))
			records.push(record)
	}).then(() => records)
}

arn.batchGet = Promise.promisify(function(set, keys, callback) {
	arn.db.batchGet(keys.map(key => {
		return {
			ns: 'arn',
			set: set,
			key: key
		}
	}), function(error, results) {
		if(callback)
			callback(error.code !== 0 ? error : undefined, results.map(result => result.record))
		else if(error.code !== 0)
			console.error(error, error.stack)
	})
})