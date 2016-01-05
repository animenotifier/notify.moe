'use strict'

exports.post = (request, response) => {
	let user = request.user

	if(!user) {
		response.end('Not logged in')
		return
	}

	let endpoint = request.body.endpoint
	let deviceId = endpoint.split('/')[0]

	console.log('Saving device', deviceId, 'for user', user.nick)

	response.end('Subscribe response.')
}