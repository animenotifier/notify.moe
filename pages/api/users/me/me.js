'use strict'

let arn = require('../../../../lib')

exports.post = function(request, response) {
	if(request.body.function !== 'save') {
		response.end()
		return
	}

	let user = request.user

	if(!user) {
		response.end()
		return
	}

	let key = request.body.key
	let value = request.body.value

	/*if(!user.hasOwnProperty(key)) {
		response.end()
		return
	}*/

	/*if(key === 'nick') {
		let oldNick = user.nick

		arn.getAsync('NickToUser', value)
		.catch(error => {
			arn.remove('NickToUser', oldNick)
			arn.set('NickToUser', value, { userId: user.id })
		})
	}

	user[key] = value
	arn.setUser(user.id, user)*/

	response.end()
}