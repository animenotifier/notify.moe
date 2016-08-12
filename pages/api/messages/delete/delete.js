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
	
	// TODO: Check message author ID
	
	yield arn.remove('Messages', messageId)
	
	console.log(`${user.nick} deleted the message '${messageId}'`)
	response.end('success')
}