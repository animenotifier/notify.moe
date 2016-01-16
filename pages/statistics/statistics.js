'use strict'


let increment = function(obj, key) {
	if(obj.hasOwnProperty(key))
		obj[key] += 1
	else
		obj[key] = 1
}

exports.get = function(request, response) {
	let recordCount = 0
	let gender = {
		male: 0,
		female: 0,
		unknown: 0
	}
	let providers = {
		list: {},
		anime: {},
		airingDate: {}
	}

	arn.forEach('Users', function(user) {
		if(user.gender === 'male' || user.gender === 'female')
			gender[user.gender] += 1
		else
			gender.unknown += 1

		increment(providers.list, user.providers.list)
		increment(providers.anime, user.providers.anime)
		increment(providers.airingDate, user.providers.airingDate)

		recordCount++
	}).then(function() {
		response.render({
			users: {
				total: recordCount,
				gender
			},
			providers
		})
	})
}