let shortid = require('shortid')

const maxMessageLength = 100000

exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}
	
	let recipientNick = request.body.recipient
	
	if(!recipientNick) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('No recipient specified')
		return
	}
	
	let recipient = yield arn.db.getUserByNick(recipientNick)

	let text = request.body.text

	if(!text) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Message text required')
		return
	}
	
	text = text.trim()
	
	if(text.length > maxMessageLength) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Message too long')
		return
	}
	
	let messageId = shortid.generate()
	
	// Save message
	yield arn.db.set('Messages', messageId, {
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