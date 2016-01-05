'use strict'

exports.post = (request, response) => {
	let endpoint = null

	try {
		endpoint = JSON.parse(request.body).endpoint
		console.log(endpoint)
	} catch(e) {
		console.error(e)
	}

	response.end('Subscribe response.')
}