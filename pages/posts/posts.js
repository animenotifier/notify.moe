exports.get = function*(request, response) {
	let user = request.user
	let postId = request.params[0]

	if(!postId) {
		response.render({
			user
		})
		return
	}

	let post = yield arn.db.get('Posts', postId)
	yield [
		arn.db.get('Users', post.authorId).then(author => post.author = author),
		arn.db.get('Threads', post.threadId).then(thread => post.thread = thread)
	]

	// Open Graph
	request.og = {
		url: app.package.homepage + '/posts/' + postId,
		title: `${post.author.nick}'s reply to "${post.thread.title}"`,
		image: post.author.avatar ? post.author.avatar.replace('//www.gravatar.com', 'https://www.gravatar.com') : '',
		description: post.text,
		type: 'article'
	}

	response.render({
		user,
		post
	})
}