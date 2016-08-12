let shortid = require('shortid')

const maxMessageLength = 100000

exports.get = function*(request, response) {
	let user = request.user
	let viewUserNick = request.params[0]
	let viewUser = viewUserNick ? yield arn.getUserByNick(viewUserNick) : user
	
	if(!viewUser) {
		response.json({
			error: "Invalid recipient"
		})
		return
	}
	
	let messages = yield arn.filter('Messages', message => message.recipientId === viewUser.id)
	
	let idToUser = {}
	messages.forEach(message => {
		idToUser[message.authorId] = null
		idToUser[message.recipientId] = null
	})
	
	let users = yield arn.batchGet('Users', Object.keys(idToUser))
	
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
		
		delete message.authorId
		delete message.recipientId
	})
	
	response.json(messages)
}

exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}
	
	let recipientNick = request.body.recipient
	
	if(!recipientNick) {
		response.writeHead(409)
		response.end('No recipient specified')
		return
	}
	
	let recipient = yield arn.getUserByNick(recipientNick)

	let text = request.body.text

	if(!text) {
		response.writeHead(409)
		response.end('Message text required')
		return
	}
	
	text = text.trim()
	
	if(text.length > maxMessageLength) {
		response.writeHead(409)
		response.end('Message too long')
		return
	}
	
	let messageId = shortid.generate()
	
	// Save message
	yield arn.set('Messages', messageId, {
		id: messageId,
		authorId: user.id,
		recipientId: recipient.id,
		text,
		likes: [],
		created: (new Date()).toISOString()
	})
	
	// Send notification about the message
	yield arn.sendNotification(recipient, {
		title: `New message from ${user.nick}`,
		icon: user.avatar + '?s=128&r=x&d=mm',
		body: text
	})
	
	response.end('success')
}