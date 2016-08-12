exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}
	
	let messageId = request.params[0]
	
	if(!messageId) {
		response.writeHead(409)
		response.end('No message specified')
		return
	}
	
	let message = yield arn.get('Messages', messageId)
	
	if(!message.likes)
		message.likes = []
	
	if(message.likes.indexOf(user.id) !== -1) {
		response.end('Already liked that message')
		return
	}
	
	message.likes.push(user.id)
	
	yield arn.set('Messages', message.id, {
		likes: message.likes
	})
	
	console.log(`${user.nick} liked the message '${messageId}'`)
	response.end('success')
}