'use strict'

let arn = require('../../lib')

exports.get = function(request, response) {
	let recordCount = 0
	let gender = {
		male: 0,
		female: 0,
		unknown: 0
	}

	arn.scan('Users', function(user) {
		if(user.gender === 'male' || user.gender === 'female')
			gender[user.gender] += 1
		else
			gender.unknown += 1

		recordCount++
	}, function() {
		response.render({
			users: {
				total: recordCount,
				gender
			},
		})
	})
}