import * as arn from 'arn'
import * as HTTP from 'http-status-codes'
import * as shortid from 'shortid'

const minTitleLength = 3
const maxTitleLength = 100
const maxPostLength = 100000

exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}

	let tag = request.body.tag

	if(!tag) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('No tag specified')
		return
	}

	let title = request.body.title

	if(!title) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Thread title required')
		return
	}

	title = title.trim()

	if(title.length > maxTitleLength || title.length < minTitleLength) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Invalid title length')
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

	let sticky = request.body.sticky ? 1 : 0

	if((sticky || tag === 'update') && user.role !== 'admin') {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not authorized')
		return
	}

	let threadId = shortid.generate()

	// Save post
	yield arn.db.set('Threads', threadId, {
		id: threadId,
		authorId: user.id,
		title,
		text,
		tags: [tag],
		likes: [],
		sticky,
		replies: 0,
		created: (new Date()).toISOString()
	})

	response.end(threadId)

	// Announce on chat
	// arn.chatBot.sendMessage('forum', `New thread: ${app.package.homepage}/threads/${threadId}`)
}