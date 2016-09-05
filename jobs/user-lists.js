let updateUserLists = coroutine(function*() {
	let tasks = []
	let allUsers = yield arn.filter('Users', user => arn.isActiveUser(user))

	console.log(`${allUsers.length} active users`)

	for(let orderBy of Object.keys(arn.userOrderBy)) {
		let method = arn.userOrderBy[orderBy]
		let categories = method.getCategories()
		let addUser = method.addUser
		let cacheKey = `users:${orderBy}`

		console.log(chalk.yellow('✖'), `Updating user list ${chalk.yellow(orderBy)}...`)

		let updateUserList = coroutine(function*() {
			let users = yield Promise.filter(allUsers, user => addUser(user, categories))

			// Sort by registration date
			Object.keys(categories).forEach(categoryName => {
				let category = categories[categoryName]
				category.sort((a, b) => {
					let aVatar = a.avatar !== ''
					let bVatar = b.avatar !== ''

					if(aVatar === bVatar)
						return new Date(b.lastView.date) - new Date(a.lastView.date)
					else
						return bVatar - aVatar
				})

				// Reduce data size for the database
				categories[categoryName] = category.map(user => {
					return {
						id: user.id,
						nick: user.nick,
						avatar: user.avatar
					}
				})
			})

			console.log(chalk.green('✔'), `Finished updating user list ${chalk.yellow(orderBy)}`)

			return arn.set('Cache', cacheKey, {
				categories
			}).catch(error => {
				console.error(`Error saving user list ${chalk.yellow(orderBy)}`, error)
			})
		})

		yield updateUserList()
		yield Promise.delay(100)
	}

	yield tasks

	console.log(chalk.green('✔'), 'Finished updating all user lists')
})

arn.repeatedly(4 * 60, updateUserLists)