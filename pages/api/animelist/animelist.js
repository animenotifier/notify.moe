'use strict'



exports.get = function(request, response) {
	let nick = request.params[0]
	let command = request.params[1]
	let clearCache = false

	if(!nick) {
		response.writeHead(409)
		return response.json({
			error: 'Username not specified'
		})
	}

	if(command && command === 'clearListCache') {
		clearCache = true
	}

	return arn.getAnimeListByNick(nick, clearCache).then(json => {
		// Delete critical data
		delete json.userId

		response.json(json)
	}).catch(error => {
		response.writeHead(409)
		response.json({
			error: error.toString()
		})
	})
}