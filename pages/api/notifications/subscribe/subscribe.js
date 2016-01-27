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

	let parts = endpoint.split('/')
	let deviceId = parts[parts.length - 1]
	let pushUrl = endpoint.slice(-deviceId.length)

	console.log('Saving push endpoint', endpoint, 'for user', user.nick)

	// Add endpoint
	user.pushEndpoints[endpoint] = {
		registered: (new Date()).toISOString()
	}

	arn.set('Users', user.id, user).then(() => response.end('success'))
}