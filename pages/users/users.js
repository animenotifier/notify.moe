'use strict'

let arn = require('../../lib')
let gravatar = require('gravatar')
let NodeCache = require('node-cache')

let cache = new NodeCache({
	stdTTL: 5 * 60
})

exports.get = function(request, response) {
	let orderBy = request.params[0]
	let addUser = null
	let now = new Date()
	let cacheKey = `users:${orderBy}`

	cache.get(cacheKey, function(error, categories) {
		if(!error && categories) {
			response.render({
				categories
			})
			return
		}

		categories = {}

		switch(orderBy) {
			case 'countries':
				addUser = user => {
					if(!user.location)
						return

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
					if(categories.hasOwnProperty(user.providers.list))
						categories[user.providers.list].push(user)
					else
						categories[user.providers.list] = [user]
				}
				break

			case 'login':
			case 'registration':
				// Force a special order
				categories = {
					'Last 24 hours': [],
					'Yesterday': [],
					'This week': [],
					'This month': [],
					'Ojii-san': []
				}

				addUser = user => {
					let date = new Date(orderBy === 'registration' ? user.registered : user.lastLogin)
					let seconds = Math.floor((now - date) / 1000)
					let days = seconds / 60 / 60 / 24
					let categoryName = 'Ojii-san'

					if(days <= 1)
						categoryName = 'Last 24 hours'
					else if(days <= 2)
						categoryName = 'Yesterday'
					else if(days <= 7)
						categoryName = 'This week'
					else if(days <= 30)
						categoryName = 'This month'

					if(categories.hasOwnProperty(categoryName))
						categories[categoryName].push(user)
					else
						categories[categoryName] = [user]
				}
				break

			default:
				categories.All = []
				addUser = user => categories.All.push(user)
		}

		arn.scan('Users', function(user) {
			if(!arn.isActiveUser(user))
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

			cache.set(cacheKey, categories, () => {})
		})
	})
}