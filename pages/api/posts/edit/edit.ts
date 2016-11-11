import * as arn from 'arn'
import * as HTTP from 'http-status-codes'

const parameters = ['id', 'text']
const maxPostLength = 100000

exports.post = async function(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}

	if(!arn.assertParams(request, response, parameters))
		return

	let postId = request.body.id
	let text = request.body.text

	text = text.trim()

	if(text.length > maxPostLength) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Post too long')
		return
	}

	let post = await arn.db.get('Posts', postId)

	if(post.authorId !== user.id) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Can not edit the post of a different user')
		return
	}

	// Save post
	await arn.db.set('Posts', postId, {
		text,
		edited: (new Date()).toISOString()
	})

	response.end(this.app.markdown(text))
}