'use strict'

exports.get = function(request, response) {
	let user = request.user

	response.render({
		user,
		widgets: arn.news
	})
}