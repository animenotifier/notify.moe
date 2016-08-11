'use strict'

const maxMessages = 5

let gravatar = require('gravatar')

exports.get = function*(request, response) {
	let user = request.user
	let viewUserNick = request.params[0]
	
	/*let messages = [
		{
			gravatar: '//www.gravatar.com/avatar/c53fee0670aefc8e592465f598d10a83?s=50',
			text: '# Test\nThis is a test message. This is a test message. This is a test message. This is a test message. This is a test message.\n\n* Point 1\n* Point 2\n* Point 3\n\nI just wanted to say hi. Here\'s a [link](http://google.com).'
		},
		{
			gravatar: '//www.gravatar.com/avatar/794b90f30252319f6a6ff8eeb0d4ecd7?s=50',
			text: 'This is a 2nd test message.'
		},
		{
			gravatar: '//www.gravatar.com/avatar/b56f02c9addef85dd82ec7bb7a21173d?s=50',
			text: 'This is a 3rd test message.'
		}
	]*/

	try {
		let viewUser = yield arn.getUserByNick(viewUserNick)
		let messages = yield arn.filter('Messages', message => message.recipientId === viewUser.id)
		
		messages.sort((a, b) => a.created < b.created)
		
		if(messages.length > maxMessages)
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