const maxMessages = 5

let gravatar = require('gravatar')

exports.get = function*(request, response) {
	let user = request.user
	
	let viewUser = null
	let viewUserNick = request.params[0]
	
	if(!viewUserNick) {
		viewUser = user
	} else {
		viewUser = yield arn.db.getUserByNick(viewUserNick)
	}
	
	if(!viewUser) {
		response.render({
			user
		})
		return
	}

	try {
		let messages = yield arn.db.filter('Messages', message => message.recipientId === viewUser.id)
		
		messages.sort((a, b) => (a.created > b.created) ? -1 : ((a.created < b.created) ? 1 : 0))
		
		if(viewUserNick && messages.length > maxMessages)
			messages.length = maxMessages
		
		yield messages.map(message => {
			return arn.db.get('Users', message.authorId).then(author => message.author = author)
		})

		response.render({
			user,
			viewUser,
			messages
		})
	} catch(error) {
		console.error(error)
		response.render({
			user
		})
	}
}