exports.get = function(request, response) {
	let nick = request.params[0]

	if(!nick)
		return response.end()

	arn.getUserByNick(nick).then(user => {
		user.notificationsEnabled = Object.keys(user.pushEndpoints).length > 0

		// Do not show critical information
		if(!request.user || request.user.role !== 'admin') {
			delete user.id
			delete user.ip
			delete user.accounts
			delete user.ageRange
			delete user.email
			delete user.gender
			delete user.firstName
			delete user.lastName
			delete user.agent
			delete user.location
			delete user.lastLogin
			delete user.pushEndpoints
		}

		response.json(user)
	}).catch(error => {
		response.json({
			error
		})
	})
}