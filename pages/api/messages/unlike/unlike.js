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
	let index = message.likes.indexOf(user.id)
	
	if(index === -1) {
		response.end('You did not like this message yet')
		return
	}
	
	message.likes.splice(index, 1)
	
	yield arn.set('Messages', message.id, {
		likes: message.likes
	})
	
	console.log(`${user.nick} unliked the message '${messageId}'`)
	response.end('success')
}