let updateUserLists = coroutine(function*() {
	let tasks = []
	
	for(let orderBy of Object.keys(arn.userOrderBy)) {
		console.log(chalk.yellow('✖'), `Updating user list ${chalk.yellow(orderBy)}...`)

		let method = arn.userOrderBy[orderBy]
		let categories = method.getCategories()
		let addUser = method.addUser
		let cacheKey = `users:${orderBy}`

		tasks.push(arn.filter('Users', user => arn.isActiveUser(user) && user.avatar).then(coroutine(function*(users) {
			users = yield Promise.filter(users, user => addUser(user, categories))

			// Sort by registration date
			Object.keys(categories).forEach(categoryName => {
				let category = categories[categoryName]
				category.sort((a, b) => new Date(a.registered) - new Date(b.registered))
				
				// Reduce data size for the database
				categories[categoryName] = category.map(user => {
					return {
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
		})))
		
		tasks.push(Promise.delay(1000))
	}
	
	yield tasks
})

arn.repeatedly(4 * 60, updateUserLists)