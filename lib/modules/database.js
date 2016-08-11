'use strict'

let Promise = require('bluebird')
let Aerospike = require('aerospike')
let Key = Aerospike.Key

arn.db = Aerospike.client({
    hosts: [{
        addr: '127.0.0.1',
        port: 3000
    }],
    log: {
        level: Aerospike.log.WARN
    },
    policies: {
        timeout: 5000
    },
	maxConnsPerNode: 1024
})

arn.connectDB = Promise.promisify(callback => {
	arn.db.connect(error => {
	    if(error) {
			console.error('Couldn\'t connect to database!')
			console.error(error)

			if(callback)
				callback(error)
	    } else {
			console.log('Successfully connected to database!')

			arn.events.emit('database ready')

			if(callback)
				callback(undefined)
	    }
	})
})

arn.db.ready = arn.connectDB()

arn.get = Promise.promisify(function(set, key, callback) {
	arn.db.get(new Key('arn', set, key), function(error, record, metadata) {
		if(callback)
			callback(error, record)
		else if(error && error.code !== 0)
			console.error(error, error.stack)
	})
})

arn.set = Promise.promisify(function(set, key, obj, callback) {
	arn.db.put(new Key('arn', set, key), obj, function(error) {
		if(callback)
			callback(error, obj)
		else if(error.code !== 0)
			console.error(error, error.stack)
	})
})

arn.remove = Promise.promisify(function(set, key, callback) {
	arn.db.remove(new Key('arn', set, key), function(error, key) {
		if(callback)
			callback(error, key)
		else if(error.code !== 0)
			console.error(error, error.stack)
	})
})

arn.forEach = Promise.promisify(function(set, func, callback) {
	let scan = arn.db.scan('arn', set)
	scan.concurrent = true
	scan.nobins = false
	
	if(arn.runningBackgroundJobs)
		scan.priority = Aerospike.scanPriority.LOW
	else
		scan.priority = Aerospike.scanPriority.HIGH

	let stream = scan.foreach()

	stream.on('data', func)

	stream.on('error', function(error) {
		console.error('Error occured while scanning:', error)
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
	arn.db.batchRead(keys.map(key => {
		return {
			key: new Key('arn', set, key),
			read_all_bins: true
		}
	}), function(error, results) {
		if(callback)
			callback(error, results.map(result => result.bins))
		else if(error.code !== 0)
			console.error(error, error.stack)
	})
})