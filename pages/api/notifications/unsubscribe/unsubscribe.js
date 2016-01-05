'use strict'

exports.post = (request, response) => {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}

	console.log(request.body)
	console.log(request.body.deviceId)

	response.end('test')
}