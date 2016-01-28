'use strict'

let gravatar = require('gravatar')

let getAddUserFunction = (orderBy) => {
	return (user, categories) => {
		let date = new Date(orderBy === 'joindate' ? user.registered : user.lastLogin)
		let now = new Date()
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
}

let orderByMethods = {
	countries: {
		getCategories: () => {
			return {}
		},

		addUser: (user, categories) => {
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
	},

	'providers/list': {
		getCategories: () => {
			return {
				'AniList': [],
				'HummingBird': [],
				'MyAnimeList': [],
				'AnimePlanet': []
			}
		},

		addUser: (user, categories) => {
			if(categories.hasOwnProperty(user.providers.list))
				categories[user.providers.list].push(user)
			else
				categories[user.providers.list] = [user]

			return true
		}
	},

	'providers/anime': {
		getCategories: () => {
			return {
				'CrunchyRoll': [],
				'Nyaa': []
			}
		},

		addUser: (user, categories) => {
			if(categories.hasOwnProperty(user.providers.anime))
				categories[user.providers.anime].push(user)
			else
				categories[user.providers.anime] = [user]

			return true
		}
	},

	login: {
		getCategories: () => {
			return {
				'Last 24 hours': [],
				'Yesterday': [],
				'This week': [],
				'This month': [],
				'Ojii-san': []
			}
		},

		addUser: getAddUserFunction('login')
	},

	joindate: {
		getCategories: () => {
			return {
				'Last 24 hours': [],
				'Yesterday': [],
				'This week': [],
				'This month': [],
				'Ojii-san': []
			}
		},

		addUser: getAddUserFunction('joindate')
	},

	osu: {
		getCategories: () => {
			return {
				'10k pp': [],
				'9k pp': [],
				'8k pp': [],
				'7k pp': [],
				'6k pp': [],
				'5k pp': [],
				'4k pp': [],
				'3k pp': [],
				'2k pp': [],
				'1k pp': [],
				'Beginners': [],
			}
		},

		addUser: (user, categories) => {
			if(!user.osu || !user.osuDetails)
				return false

			if(user.osuDetails.pp < 1000)
				categories.Beginners.push(user)
			else if(user.osuDetails.pp >= 10000)
				categories['10k pp'].push(user)
			else
				categories[parseInt(user.osuDetails.pp / 1000) + 'k pp'].push(user)

			return true
		}
	},

	staff: {
		getCategories: () => {
			return {
				'Developers': [],
				'Editors': []
			}
		},

		addUser: (user, categories) => {
			if(user.role === 'admin') {
				categories.Developers.push(user)
				return true
			} else if(user.role === 'editor') {
				categories.Editors.push(user)
				return true
			}

			return false
		}
	},

	default: {
		getCategories: () => {
			return {
				All: []
			}
		},

		addUser: (user, categories) => categories.All.push(user)
	}
}

arn.repeatedly(5 * 60, () => {
	Object.keys(orderByMethods).forEach(orderBy => {
		arn.cacheLimiter.removeTokens(1, () => {
			let method = orderByMethods[orderBy]
			let categories = method.getCategories()
			let addUser = method.addUser
			let cacheKey = `users:${orderBy}`

			arn.filter('Users', user => arn.isActiveUser(user) && addUser(user, categories)).then(users => {
				users.forEach(user => {
					user.gravatarURL = gravatar.url(user.email, {s: '50', r: 'x', d: '404'}, true)
				})

				// Sort by registration date
				Object.keys(categories).forEach(categoryName => {
					let category = categories[categoryName]
					category.sort((a, b) => new Date(a.registered) - new Date(b.registered))
				})

				return arn.set('Cache', cacheKey, {
					categories
				})
			})
		})
	})
})

exports.get = function(request, response) {
	let orderBy = request.params[0] || 'default'

	if(request.params[1])
		orderBy += '/' + request.params[1]

	let cacheKey = `users:${orderBy}`

	arn.get('Cache', cacheKey).then(record => {
		// We need to copy the object to keep the order.
		// It sucks but the database doesn't keep object properties' order.
		let categories = orderByMethods[orderBy].getCategories()

		Object.keys(record.categories).forEach(categoryName => {
			categories[categoryName] = record.categories[categoryName]
		})

		response.render({
			categories
		})
	}).catch(error => {
		response.render({})
	})
}