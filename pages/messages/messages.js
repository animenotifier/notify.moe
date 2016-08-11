'use strict'

const maxMessages = 5

let gravatar = require('gravatar')

exports.get = function*(request, response) {
	let user = request.user
	
	let viewUser = null
	let viewUserNick = request.params[0]
	
	if(!viewUserNick) {
		viewUser = user
	} else {
		viewUser = yield arn.getUserByNick(viewUserNick)
	}
	
	if(!viewUser) {
		response.render({
			user
		})
		return
	}

	try {
		let messages = yield arn.filter('Messages', message => message.recipientId === viewUser.id)
		
		messages.sort((a, b) => a.created < b.created)
		
		if(viewUserNick && messages.length > maxMessages)
			messages.length = maxMessages
		
		yield messages.map(message => {
			return arn.get('Users', message.authorId).then(author => message.author = author)
		})

		response.render({
			user,
			viewUser,
			messages
		})
	} catch(error) {
		console.error(error, error.stack)
		response.render({
			user
		})
	}
}