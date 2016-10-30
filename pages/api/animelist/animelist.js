exports.get = function(request, response) {
	let nick = request.params[0]
	let command = request.params[1]
	let clearCache = false

	if(!nick) {
		response.writeHead(HTTP.BAD_REQUEST)
		return response.json({
			error: 'Username not specified'
		})
	}

	if(command && (command === 'clearCache' || command === 'clearListCache')) {
		clearCache = true
	}

	return arn.getAnimeListByNick(nick, clearCache).then(json => {
		// Delete critical data
		delete json.userId

		response.json(json)
	}).catch(error => {
		response.writeHead(HTTP.BAD_REQUEST)
		response.json({
			error: error.toString(),
			stack: error.stack ? error.stack.toString().replace(/\/home\/eduard\//g, '').split('\n') : ''
		})
	})
}