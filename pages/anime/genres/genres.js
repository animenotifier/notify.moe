'use strict'

exports.get = (request, response) => {
	let genre = request.params[0]

	if(!genre) {
		response.render({})
		return
	}

	arn.get('Genres', genre).then(record => {
		response.render(record)
	}).catch(error => {
		console.error(error, error.stack)
		response.render({})
	})
}