'use strict'

let arn = require('../../../../lib')

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

	// Add ID to the list of devices
	user.devices.push(deviceId)
	
	arn.setUserAsync(user.id, user).then(() => response.end('Subscribe response.'))
}