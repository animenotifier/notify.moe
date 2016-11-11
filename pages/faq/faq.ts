exports.get = (request, response) => {
	response.render({
		user: request.user
	})
}