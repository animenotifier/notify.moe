const maxPostLength = 100000

exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}

	let text = request.body.text

	if(!text) {
		response.writeHead(409)
		response.end('Post text required')
		return
	}
	
	text = text.trim()
	
	if(text.length > maxPostLength) {
		response.writeHead(409)
		response.end('Post too long')
		return
	}
	
	let postId = request.body.id
	
	if(!postId) {
		response.writeHead(409)
		response.end('Post ID required')
		return
	}
	
	let post = yield arn.get('Posts', postId)
	
	if(post.authorId !== user.id) {
		response.writeHead(409)
		response.end('Can not edit the post of a different user')
		return
	}
	
	// Save post
	yield arn.set('Posts', postId, {
		text,
		edited: (new Date()).toISOString()
	})
	
	response.end(this.app.markdown(text))
}