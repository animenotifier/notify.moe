const maxThreadCount = 25

exports.get = function*(request, response) {
	let user = request.user
	let threads = yield arn.all('Threads')
	
	if(threads.length > maxThreadCount)
		threads.length = maxThreadCount
		
	let users = yield arn.batchGet('Users', threads.map(thread => thread.authorId))
	let idToUser = {}
	
	users.forEach(user => idToUser[user.id] = user)
	threads.forEach(thread => thread.author = idToUser[thread.authorId])
	
	response.render({
		user,
		threads,
		idToUser
	})
}