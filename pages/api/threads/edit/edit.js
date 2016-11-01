const parameters = ['id', 'text']
const maxThreadLength = 100000

exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}

	if(!arn.assertParams(request, response, parameters))
		return

	let threadId = request.body.id
	let text = request.body.text

	text = text.trim()

	if(text.length > maxThreadLength) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Thread too long')
		return
	}

	let thread = yield arn.db.get('Threads', threadId)

	if(thread.authorId !== user.id) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Can not edit the thread of a different user')
		return
	}

	// Save post
	yield arn.db.set('Threads', threadId, {
		text,
		edited: (new Date()).toISOString()
	})

	response.end(this.app.markdown(text))
}