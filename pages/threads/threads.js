exports.get = function*(request, response) {
	let user = request.user
	let threadId = request.params[0]
	
	if(!threadId) {
		response.render({
			user
		})
		return
	}
	
	let thread = yield arn.get('Threads', threadId)
	yield arn.get('Users', thread.authorId).then(author => thread.author = author)
	
	response.render({
		user,
		thread
	})
}