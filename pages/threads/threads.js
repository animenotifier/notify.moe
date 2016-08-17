exports.get = function*(request, response) {
	let user = request.user
	let threadId = request.params[0]
	
	if(!threadId) {
		response.render({
			user
		})
		return
	}
	
	let thread = yield arn.get('Threads', threadId)
	yield arn.get('Users', thread.authorId).then(author => thread.author = author)
	
	let posts = yield arn.filter('Posts', post => post.threadId === threadId)
	
	// const testTexts = [
	// 	'Lorem ipsum dolor sit amet',
	// 	'Lorem ipsum',
	// 	'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus at suscipit enim. Proin nec magna non lacus molestie sollicitudin ornare eu risus. Praesent tincidunt sapien at est convallis, sit amet scelerisque nisl ultrices. Proin vitae turpis semper, efficitur neque et, pellentesque mauris. In dui nisl, elementum ultrices mollis a, pretium et enim. Quisque eget quam quis mauris ornare rutrum. Aliquam finibus metus eget magna dapibus, non porta nibh venenatis. Mauris eget ullamcorper lorem, id blandit magna. Integer non dui sed diam consequat viverra.',
	// 	'Lorem ipsum dolor sit amet, consectetur adipiscing elit.',
	// 	'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus at suscipit enim. Proin nec magna non lacus molestie sollicitudin ornare eu risus. Praesent tincidunt sapien at est convallis, sit amet scelerisque nisl ultrices. Proin vitae turpis semper, efficitur neque et, pellentesque mauris. In dui nisl, elementum ultrices mollis a, pretium et enim. Quisque eget quam quis mauris ornare rutrum. Aliquam finibus metus eget magna dapibus, non porta nibh venenatis. Mauris eget ullamcorper lorem, id blandit magna. Integer non dui sed diam consequat viverra.\n\nDonec porttitor elementum luctus. Fusce luctus, justo nec dictum rutrum, risus tortor maximus purus, finibus ornare enim erat ut urna. Nulla imperdiet quam nec sapien viverra, id maximus mi commodo. Maecenas ac vestibulum dolor, sit amet aliquam lectus. Integer ut nulla dapibus, porta ipsum et, pharetra mi. Suspendisse cursus, metus eu tempor venenatis, est augue varius eros, congue hendrerit tortor sapien vel elit. Etiam ornare mi eu neque molestie, ut mollis purus accumsan. Nunc vitae feugiat felis. Vivamus magna orci, molestie a justo at, finibus accumsan erat. Etiam hendrerit felis quis ligula aliquet venenatis. Nulla placerat, enim a fringilla tincidunt, felis nibh egestas risus, at elementum quam magna eget nunc. Maecenas nec sodales nisl. Vestibulum sollicitudin arcu id nunc gravida commodo. Lorem ipsum dolor sit amet, consectetur adipiscing elit.'
	// ]
	// 
	// for(let i = 0; i < 150; i++) {
	// 	posts.push({
	// 		id: 'post' + i,
	// 		threadId,
	// 		authorId: ['HyhW-TsW', 'VkBaMJ6ux', 'NyGyZ2xwe', 'EkffWKXte'][i % 4],
	// 		text: testTexts[i % testTexts.length],
	// 		likes: [],
	// 		created: (new Date()).toISOString()
	// 	})
	// }
	// 
	// yield posts.map(post => arn.set('Posts', post.id, post))
	// console.log('saved')
	
	posts.sort((a, b) => {
		return (a.created > b.created) ? 1 : ((a.created < b.created) ? -1 : 0)
	})
	
	let users = yield arn.batchGet('Users', posts.map(post => post.authorId))
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