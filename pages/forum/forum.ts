import * as arn from 'arn'
import { User } from 'arn/interfaces/User'

const maxThreadCount = 12
const tagToIcon = {
	general: 'paperclip',
	news: 'newspaper-o',
	anime: 'television',
	update: 'cubes',
	suggestion: 'lightbulb-o',
	bug: 'bug'
}
const openGraph = {
	url: app.package.homepage + '/forum',
	title: 'Anime Notifier - Forum',
	description: 'Forum for notify.moe'
}

exports.get = async function(request, response) {
	let user = request.user
	let tag = request.params[0]
	let threads = null

	if(!tag)
		threads = await arn.db.all('Threads')
	else
		threads = await arn.db.filter('Threads', thread => thread.tags && thread.tags.indexOf(tag) !== -1)

	threads.sort((a, b) => {
		let order = (a.sticky ? (-1 + b.sticky) : b.sticky)

		if(order === 0)
			return (a.created > b.created) ? -1 : ((a.created < b.created) ? 1 : 0)

		return order
	})

	await Promise.all(threads.map(thread => arn.db.set('Threads', thread.id, thread)))

	if(threads.length > maxThreadCount)
		threads.length = maxThreadCount

	let users: Array<User> = await arn.db.getMany('Users', threads.map(thread => thread.authorId))
	let idToUser = {}

	users.forEach(user => idToUser[user.id] = user)
	threads.forEach(thread => {
		thread.author = idToUser[thread.authorId]

		if(!thread.tags)
			thread.icons = []
		else
			thread.icons = thread.tags.map(tag => tagToIcon[tag]).filter(icon => icon)
	})

	// Open Graph
	request.og = openGraph

	response.render({
		user,
		threads,
		idToUser
	})
}