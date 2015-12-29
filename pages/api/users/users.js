'use strict'

let arn = require('../../../lib')

module.exports = {
	get: function(request, response) {
		let nick = request.params[0]

		if(!nick)
			return response.end()

		arn.getUserByNick(nick, function(error, user) {
			if(error)
				return response.end()

			response.setHeader('Content-Type', 'application/json')

			// Do not show critical information
			delete user.id
			delete user.ip
			delete user.accounts
			delete user.ageRange
			delete user.email
			delete user.gender
			delete user.firstName
			delete user.lastName

			response.end(JSON.stringify(user))
		})
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
}