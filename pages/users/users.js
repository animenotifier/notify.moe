exports.get = function(request, response) {
	let orderBy = request.params[0] || 'default'

	if(request.params[1])
		orderBy += '/' + request.params[1]

	let cacheKey = `users:${orderBy}`

	arn.db.get('Cache', cacheKey).then(record => {
		// We need to copy the object to keep the order.
		// It sucks but the database doesn't keep object properties' order.
		let categories = arn.userOrderBy[orderBy].getCategories()

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