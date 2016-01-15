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

	console.log('Saving device', deviceId, 'for user', user.nick)

	// Add ID to the user's devices
	user.devices[deviceId] = (new Date()).toISOString()

	arn.set('Users', user.id, user).then(() => response.end('success'))
}