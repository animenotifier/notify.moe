'use strict'

let aero = require('aero')
let gravatar = require('gravatar')

exports.get = function(request, response) {
	let users = []

	let scan = aero.db.query('arn', 'Accounts', {
	    concurrent: true,
	    nobins: false
	})

	let stream = scan.execute()

	stream.on('data', function(record) {
		let user = record.bins
		user.gravatarURL = gravatar.url(user.email, {s: '50', r: 'x', d: 'mm'}, true)
		users.push(user)
	})

	stream.on('error', function(error) {
		console.error('Error occured while scanning:', error)
	})

	stream.on('end', function(end) {
		response.render({
			users
		})
	})
}