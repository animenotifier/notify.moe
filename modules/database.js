'use strict'

let Promise = require('bluebird')
let aerospike = require('aerospike')
let arn = require('../lib')

module.exports = function(aero, callback) {
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

	Promise.promisifyAll(arn.db)

	arn.db.connect(function(response) {
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
}