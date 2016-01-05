'use strict'

exports.post = (request, response) => {
	let endpoint = null

	try {
		endpoint = JSON.parse(request.body).endpoint
	} catch(e) {
		console.error(e)
	}
	
	response.end('Unsubscribe response.')
}