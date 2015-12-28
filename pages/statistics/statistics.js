'use strict'

let aero = require('aero')

exports.get = function(request, response) {
	let scan = aero.db.query('arn', 'Accounts', {
	    concurrent: true,
	    nobins: false,
	})
	let stream = scan.execute()
	let recordCount = 0

	stream.on('data', function(record) {
		recordCount++
	})

	stream.on('error', function(error) {
		console.error('Error occured while scanning:', error)
	})

	stream.on('end', function(end) {
		response.render({
			accounts: {
				total: recordCount
			}
		})
	})
}