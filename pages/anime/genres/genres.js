exports.get = (request, response) => {
	let user = request.user
	let genre = request.params[0]

	if(!genre) {
		response.render({
			user
		})
		return
	}

	arn.db.get('Genres', genre).then(record => {
		response.render(Object.assign({
			user
		}, record))
	}).catch(error => {
		console.error(error)
		response.render({
			user
		})
	})
}