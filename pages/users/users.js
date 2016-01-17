'use strict'

let gravatar = require('gravatar')

exports.get = function(request, response) {
	let orderBy = request.params[0]
	let addUser = null
	let now = new Date()
	let cacheKey = `users:${orderBy}`

	arn.userListCache.get(cacheKey, function(error, categories) {
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
						return false

					let country = user.location.countryName

					if(!country || country === '-')
						return false

					if(categories.hasOwnProperty(country))
						categories[country].push(user)
					else
						categories[country] = [user]

					return true
				}
				break

			case 'listproviders':
				addUser = user => {
					if(categories.hasOwnProperty(user.providers.list))
						categories[user.providers.list].push(user)
					else
						categories[user.providers.list] = [user]

					return true
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

					return true
				}
				break

			default:
				categories.All = []
				addUser = user => categories.All.push(user)
		}

		arn.filter('Users', user => arn.isActiveUser(user) && addUser(user)).then(users => {
			users.forEach(user => {
				user.gravatarURL = gravatar.url(user.email, {s: '50', r: 'x', d: '404'}, true)
			})

			// Sort by registration date
			Object.keys(categories).forEach(categoryName => {
				let category = categories[categoryName]
				category.sort((a, b) => new Date(a.registered) - new Date(b.registered))
			})

			response.render({
				categories
			})

			arn.userListCache.set(cacheKey, categories, () => {})
		})
	})
}