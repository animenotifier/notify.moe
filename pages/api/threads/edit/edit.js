const maxThreadLength = 100000

exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}

	let text = request.body.text

	if(!text) {
		response.writeHead(409)
		response.end('Thread text required')
		return
	}
	
	text = text.trim()
	
	if(text.length > maxThreadLength) {
		response.writeHead(409)
		response.end('Thread too long')
		return
	}
	
	let threadId = request.body.id
	
	if(!threadId) {
		response.writeHead(409)
		response.end('Thread ID required')
		return
	}
	
	// Save post
	yield arn.set('Threads', threadId, {
		text,
		edited: (new Date()).toISOString()
	})
	
	response.end(this.app.markdown(text))
}