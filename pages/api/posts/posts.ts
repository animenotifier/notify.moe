import * as arn from 'arn'
import * as HTTP from 'http-status-codes'
import * as shortid from 'shortid'

const maxPostLength = 100000

exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}

	let threadId = request.body.threadId

	if(!threadId) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('No thread specified')
		return
	}

	let thread = yield arn.db.get('Threads', threadId)

	if(!thread) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Thread does not exist')
		return
	}

	let text = request.body.text

	if(!text) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Post text required')
		return
	}

	text = text.trim()

	if(text.length > maxPostLength) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Post too long')
		return
	}

	let postId = shortid.generate()

	// Save post
	yield arn.db.set('Posts', postId, {
		id: postId,
		authorId: user.id,
		threadId: thread.id,
		text,
		likes: [],
		created: (new Date()).toISOString()
	})

	// Update reply count
	if(!thread.replies)
		thread.replies = 0

	arn.db.set('Threads', thread.id, {
		replies: thread.replies + 1
	})

	response.end(postId)

	// Announce on chat
	// arn.chatBot.sendMessage('forum', `${app.package.homepage}/posts/${postId}`)
}