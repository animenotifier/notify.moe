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

	arn.get('NickToUser', viewUserNick, (error, nick) => {
		if(error) {
			response.render({ user })
			return
		}

		arn.getUser(nick.userId, function(error, viewUser) {
			if(error) {
				response.render({ user })
				return
			}

			viewUser.gravatarURL = gravatar.url(viewUser.email, {s: '200', r: 'x', d: 'mm'}, true)
			response.render({
				user,
				viewUser
			})
		})
	})
}