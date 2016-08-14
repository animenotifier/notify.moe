const parameters = ['id', 'text']
const maxThreadLength = 100000

exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}
	
	if(!arn.assertParams(request, parameters))
		return
	
	let threadId = request.body.id
	let text = request.body.text
	
	text = text.trim()
	
	if(text.length > maxThreadLength) {
		response.writeHead(409)
		response.end('Thread too long')
		return
	}
	
	let thread = yield arn.get('Threads', threadId)
	
	if(thread.authorId !== user.id) {
		response.writeHead(409)
		response.end('Can not edit the thread of a different user')
		return
	}
	
	// Save post
	yield arn.set('Threads', threadId, {
		text,
		edited: (new Date()).toISOString()
	})
	
	response.end(this.app.markdown(text))
}