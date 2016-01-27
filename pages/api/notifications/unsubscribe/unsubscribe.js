'use strict'

exports.post = (request, response) => {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}

	let endpoint = request.body.endpoint

	if(!endpoint) {
		response.writeHead(409)
		response.end('endpoint required')
		return
	}

	console.log('Removing push endpoint', endpoint, 'from user', user.nick)

	// Add ID to the user's devices
	delete user.pushEndpoints[endpoint]

	arn.set('Users', user.id, user).then(() => response.end('success'))
}