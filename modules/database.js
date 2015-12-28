'use strict'

let Promise = require('bluebird')
let aerospike = require('aerospike')

module.exports = function(aero) {
	aero.db = aerospike.client({
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

	Promise.promisifyAll(aero.db)

	aero.db.connect(function(response) {
	    if(response.code === 0) {
	        console.log('Successfully connected to database!')
	    } else {
	        console.error('Couldn\'t connect to database!')
	        console.error(response)
	    }
	})
}