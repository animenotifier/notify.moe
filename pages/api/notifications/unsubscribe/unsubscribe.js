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

	console.log('Removing push endpoint', endpoint, 'from user', user.nick)

	// Add ID to the user's devices
	delete user.pushEndpoints[endpoint]

	arn.db.set('Users', user.id, user).then(() => response.end('success'))
}