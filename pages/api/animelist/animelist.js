'use strict'

let arn = require('../../../lib')

exports.get = function(request, response) {
	let nick = request.params[0]

	if(!nick) {
		response.writeHead(409)
		return response.json({
			error: 'Username not specified'
		})
	}

	return arn.getAnimeListByNickAsync(nick).then(json => {
		response.json(json)
	}).catch(error => {
		response.writeHead(409)
		response.json({
			error: error.toString()
		})
	})
}