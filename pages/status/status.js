exports.get = function*(request, response) {
	let status = yield arn.get('Cache', 'status')
	
	response.render({
		user: request.user,
		status
	})
}