exports.get = function*(request, response) {
	let user = request.user
	let viewUserNick = request.params[0]
	let viewUser = viewUserNick ? yield arn.db.getUserByNick(viewUserNick) : user
	
	if(!viewUser) {
		response.json({
			error: "Invalid recipient"
		})
		return
	}
	
	let messages = yield arn.db.filter('Messages', message => message.recipientId === viewUser.id)
	
	let idToUser = {}
	messages.forEach(message => {
		idToUser[message.authorId] = null
		idToUser[message.recipientId] = null
	})
	
	let users = yield arn.db.getMany('Users', Object.keys(idToUser))
	
	users.forEach(user => idToUser[user.id] = user)
	messages.forEach(message => {
		let author = idToUser[message.authorId]
		let recipient = idToUser[message.recipientId]
		
		message.author = {
			nick: author.nick,
			avatar: author.avatar
		}
		
		message.recipient = {
			nick: recipient.nick,
			avatar: recipient.avatar
		}
		
		// Remove IDs
		delete message.authorId
		delete message.recipientId
		message.likes = message.likes.length
	})
	
	response.json(messages)
}