exports.get = function(request, response) {
	let user = request.user

	response.render({
		user,
		exampleName: user ? user.nick : 'YOUR_USERNAME'
	})
}