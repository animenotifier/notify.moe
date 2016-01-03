'use strict'

let arn = require('../../lib')
let gravatar = require('gravatar')

exports.get = function(request, response) {
	let orderBy = request.params[0]
	let categories = {}
	let addUser = null

	switch(orderBy) {
		case 'countries':
			addUser = user => {
				let country = user.location.countryName

				if(!country || country === '-')
					return

				if(categories.hasOwnProperty(country))
					categories[country].push(user)
				else
					categories[country] = [user]
			}
			break

		case 'listproviders':
			addUser = user => {
				let listProviderName = user.providers.list

				if(!listProviderName)
					return

				let listProvider = user.listProviders[listProviderName]

				if(!listProvider || !listProvider.userName)
					return

				if(categories.hasOwnProperty(listProviderName))
					categories[listProviderName].push(user)
				else
					categories[listProviderName] = [user]
			}
			break

		default:
			categories.Users = []
			addUser = user => categories.Users.push(user)
	}

	arn.scan('Users', function(user) {
		if(user.nick.startsWith('g') && !isNaN(parseInt(user.nick.substring(1))))
			return

		if(user.nick.startsWith('fb') && !isNaN(parseInt(user.nick.substring(2))))
			return

		user.gravatarURL = gravatar.url(user.email, {s: '50', r: 'x', d: 'mm'}, true)

		addUser(user)
	}, function() {
		// Sort by registration date
		Object.keys(categories).forEach(categoryName => {
			let category = categories[categoryName]
			category.sort((a, b) => new Date(a.registered) - new Date(b.registered))
		})

		response.render({
			categories
		})
	})
}