'use strict'

let arn = require('../../lib')

exports.get = function(request, response) {
	let user = request.user

	response.render({
		user,
		widgets: arn.news
	})
}