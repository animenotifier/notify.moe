const maxThreadCount = 25

exports.get = function*(request, response) {
	let user = request.user
	let threads = yield arn.all('Threads')
	
	if(threads.length > maxThreadCount)
		threads.length = maxThreadCount
	
	response.render({
		user,
		threads
	})
}