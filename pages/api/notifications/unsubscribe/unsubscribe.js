'use strict'

exports.post = (request, response) => {
	let endpoint = null

	try {
		endpoint = request.body.endpoint
		console.log(endpoint)
	} catch(e) {
		console.error(e)
	}

	response.end('Unsubscribe response.')
}