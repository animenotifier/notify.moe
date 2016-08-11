exports.get = function*(request, response) {
	let user = request.user
	let messageId = request.params[0]
	
	if(!messageId) {
		response.render({
			user
		})
		return
	}
	
	let message = yield arn.get('Messages', messageId)
	yield arn.get('Users', message.authorId).then(author => message.author = author)
	
	response.render({
		message,
		user
	})
}