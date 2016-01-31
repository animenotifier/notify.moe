'use strict'

let gravatar = require('gravatar')

exports.get = function*(request, response) {
	let user = request.user
	let viewUserNick = request.params[0]
	let embeddedList = request.params[1] === 'watching'

	if(!viewUserNick) {
		let viewUser = user

		if(viewUser)
			viewUser.gravatarURL = gravatar.url(viewUser.email, {s: '320', r: 'x', d: 'mm'}, true)

		response.render({
			user,
			viewUser,
			embeddedList
		})
		return
	}

	try {
		let viewUser = yield arn.getUserByNick(viewUserNick)
		viewUser.gravatarURL = gravatar.url(viewUser.email, {s: '320', r: 'x', d: 'mm'}, true)

		response.render({
			user,
			viewUser,
			embeddedList
		})
	} catch(error) {
		console.error(error, error.stack)
		response.render({
			user
		})
	}
}