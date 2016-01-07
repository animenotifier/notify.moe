'use strict'

let arn = require('../../../lib')

exports.get = function(request, response) {
	let nick = request.params[0]

	if(!nick)
		return response.end('Username not specified')

	return arn.getAnimeListAsync(nick).then(json => {
		response.json(json)
	}).catch(error => {
		console.error(error)
		response.writeHead(409)
		response.json({
			error: error.toString()
		})
	})
}