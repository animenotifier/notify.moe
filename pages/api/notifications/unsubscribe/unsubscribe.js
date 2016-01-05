'use strict'

exports.post = (request, response) => {
	let endpoint = null

	console.log('body:', request.body)

	try {
		endpoint = JSON.parse(request.body).endpoint
		console.log('json', endpoint)
	} catch(e) {
		console.error(e)
	}

	response.end('Unsubscribe response.')
}