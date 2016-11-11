import * as arn from 'arn'

exports.get = async function(request, response) {
	let user = request.user
	let threadId = request.params[0]

	if(!threadId) {
		response.render({
			user
		})
		return
	}

	let thread = await arn.db.get('Threads', threadId)
	await arn.db.get('Users', thread.authorId).then(author => thread.author = author)

	let posts = await arn.db.filter('Posts', post => post.threadId === threadId)

	posts.sort((a, b) => {
		return (a.created > b.created) ? 1 : ((a.created < b.created) ? -1 : 0)
	})

	let users = await arn.db.getMany('Users', posts.map(post => post.authorId))
	let idToUser = {}

	users.forEach(user => idToUser[user.id] = user)
	posts.forEach(post => post.author = idToUser[post.authorId])

	// Open Graph
	request.og = {
		url: app.package.homepage + '/threads/' + thread.id,
		title: thread.title,
		description: thread.text,
		type: 'article'
	}

	response.render({
		user,
		thread,
		posts
	})
}