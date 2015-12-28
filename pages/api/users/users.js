'use strict'

let arn = require('../../../lib')

module.exports = {
	get: function(request, response) {
		response.end('Welcome to the ARN users API.');
	},

	post: function(request, response) {
		let user = request.user

		if(!user) {
			response.end()
			return
		}

		if(!user.hasOwnProperty(request.body.key)) {
			response.end()
			return
		}

		if(request.body.key === 'nick') {
			let newNick = request.body.value
			let oldNick = user.nick

			arn.getAsync('NickToUser', newNick)
			.catch(error => {
				arn.remove('NickToUser', oldNick)
				arn.set('NickToUser', newNick, { userId: user.id })
			})
		}

		user[request.body.key] = request.body.value
		arn.setUser(user.id, user)

		response.end()
	}
};