exports.post = (request, response) => {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}

	let endpoint = request.body.endpoint

	if(!endpoint) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('endpoint required')
		return
	}

	let parts = endpoint.split('/')
	let deviceId = parts[parts.length - 1]
	let pushUrl = endpoint.slice(-deviceId.length)

	console.log('Saving push endpoint', endpoint, 'for user', user.nick)

	// Add endpoint
	let subscription = {
		registered: (new Date()).toISOString()
	}

	// New in Chrome 50
	if(request.body.keys)
		subscription.keys = request.body.keys

	user.pushEndpoints[endpoint] = subscription

	arn.set('Users', user.id, user).then(() => response.end('success'))
}