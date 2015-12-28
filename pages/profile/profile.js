'use strict'

let gravatar = require('gravatar')
let aero = require('aero')

exports.get = function(request, response) {
	let user = request.user
	let viewUserNick = request.params[0]

	if(!viewUserNick) {
		response.render({ user })
		return
	}

	aero.db.get({
		ns: 'arn',
		set: 'NickToAccount',
		key: viewUserNick
	}, function(error, nick, metadata, key) {
		if(error.code !== 0) {
			response.render({ user })
			return
		}

		aero.db.get({
			ns: 'arn',
			set: 'Accounts',
			key: nick.userId
		}, function(error, viewUser, metadata, key) {
			if(error.code !== 0) {
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