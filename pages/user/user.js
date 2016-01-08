'use strict'

let gravatar = require('gravatar')
let arn = require('../../lib')

exports.get = function(request, response) {
	let user = request.user
	let viewUserNick = request.params[0]

	if(!viewUserNick) {
		response.render({ user })
		return
	}

	arn.getUserByNickAsync(viewUserNick)
	.then(viewUser => {
		viewUser.gravatarURL = gravatar.url(viewUser.email, {s: '320', r: 'x', d: 'mm'}, true)
		response.render({
			user,
			viewUser
		})
	}).catch(error => {
		console.error(error.stack)
		response.render({ user })
		return
	})
}