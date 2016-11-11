exports.get = (request, response) => {
	response.render({
		user: request.user,
		tag: request.params[0] ? request.params[0] : 'general'
	})
}