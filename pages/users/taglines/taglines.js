let gravatar = require('gravatar')

exports.get = function(request, response) {
	arn.filter('Users', user => arn.isActiveUser(user) && user.tagline).then(users => {
		users.forEach(user => user.gravatarURL = gravatar.url(user.email, {s: '50', r: 'x', d: 'mm'}, true))
		users.sort((a, b) => new Date(a.registered) - new Date(b.registered))

		response.render({
			users
		})
	})
}