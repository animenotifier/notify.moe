const maxThreadCount = 13
const tagToIcon = {
	general: 'paperclip',
	anime: 'television',
	suggestion: 'lightbulb-o',
	bug: 'bug',
	news: 'newspaper-o'
}

exports.get = function*(request, response) {
	let user = request.user
	let tag = request.params[0]
	let threads = null
	
	if(!tag)
		threads = yield arn.all('Threads')
	else
		threads = yield arn.filter('Threads', thread => thread.tags && thread.tags.indexOf(tag) !== -1)
	
	// const testTexts = [
	// 	'A Guide to Forum Flags	',
	// 	'Guidelines	',
	// 	'Ban the User Above You #3 (with hashtags)	',
	// 	'Post a Random Fact About the User Above You #8 Electric Boogaloo	',
	// 	'Should There Be a Season Three of “Chuunibyou demo Koi ga Shitai”?	',
	// 	'Why would anyone vote for Trump?	',
	// 	'Count to 10k Thread V2	',
	// 	'Watched or Not Watched?	',
	// 	'Guess the Above User’s Gender	',
	// 	'Shutting Down Hummingbird	',
	// 	'Anime like Mouryou no Hako? new	',
	// 	'Any Overwatch PS4 Players? new	',
	// 	'Any anime inspired by NGE?	'
	// ]
	// 
	// for(let i = 0; i < 15; i++) {
	// 	threads.push({
	// 		title: testTexts[i % testTexts.length],
	// 		text: testTexts[i % testTexts.length],
	// 		authorId: ['HyhW-TsW', 'VkBaMJ6ux', 'NyGyZ2xwe', 'EkffWKXte'][i % 4],
	// 		id: 'test' + i,
	// 		tags: [['general', 'anime', 'suggestion', 'bug'][i % 4]],
	// 		sticky: 0,
	// 		created: (new Date()).toISOString()
	// 	})
	// }
	// 
	// threads[13].sticky = 1
	
	threads.sort((a, b) => {
		let order = (a.sticky ? (-1 + b.sticky) : b.sticky)
		
		if(order === 0)
			return (a.created > b.created) ? -1 : ((a.created < b.created) ? 1 : 0)
		
		return order
	})
	
	yield threads.map(thread => arn.set('Threads', thread.id, thread))
	
	if(threads.length > maxThreadCount)
		threads.length = maxThreadCount
		
	let users = yield arn.batchGet('Users', threads.map(thread => thread.authorId))
	let idToUser = {}
	
	users.forEach(user => idToUser[user.id] = user)
	threads.forEach(thread => {
		thread.author = idToUser[thread.authorId]
		
		if(!thread.tags)
			thread.icons = []
		else
			thread.icons = thread.tags.map(tag => tagToIcon[tag]).filter(icon => icon)
	})
	
	response.render({
		user,
		threads,
		idToUser
	})
}