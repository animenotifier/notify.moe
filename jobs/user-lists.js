'use strict'

let updateUserLists = coroutine(function*() {
	for(let orderBy of Object.keys(arn.userOrderBy)) {
		yield Promise.delay(1000)

		console.log(chalk.yellow('✖'), `Updating user list ${chalk.yellow(orderBy)}...`)

		let method = arn.userOrderBy[orderBy]
		let categories = method.getCategories()
		let addUser = method.addUser
		let cacheKey = `users:${orderBy}`

		yield arn.filter('Users', user => arn.isActiveUser(user)).then(coroutine(function*(users) {
			users = yield Promise.filter(users, user => addUser(user, categories))

			users.forEach(user => {
				user.gravatarURL = user.avatar + '?s=50'
			})

			// Sort by registration date
			Object.keys(categories).forEach(categoryName => {
				let category = categories[categoryName]
				category.sort((a, b) => new Date(a.registered) - new Date(b.registered))
			})

			return arn.set('Cache', cacheKey, {
				categories
			})
		}))

		console.log(chalk.green('✔'), `Finished updating user list ${chalk.yellow(orderBy)}`)
	}
})

arn.repeatedly(4 * 60, updateUserLists)