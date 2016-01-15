'use strict'


let gravatar = require('gravatar')

exports.get = function(request, response) {
	let users = []

	arn.scan('Users', function(user) {
		if(!arn.isActiveUser(user))
			return

		if(!user.tagline)
			return

		user.gravatarURL = gravatar.url(user.email, {s: '50', r: 'x', d: 'mm'}, true)
		users.push(user)
	}).then(function() {
		users.sort((a, b) => new Date(a.registered) - new Date(b.registered))

		response.render({
			users
		})
	})
}