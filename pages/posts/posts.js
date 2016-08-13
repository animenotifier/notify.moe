exports.get = function*(request, response) {
	let user = request.user
	let postId = request.params[0]
	
	if(!postId) {
		response.render({
			user
		})
		return
	}
	
	let post = yield arn.get('Posts', postId)
	yield arn.get('Users', post.authorId).then(author => post.author = author)
	
	response.render({
		user,
		post
	})
}