let Promise = require('bluebird')

const maxLogLength = 100

exports.get = function*(request, response) {
	if(!arn.auth(request, response, 'editor'))
		return

	let user = request.user
	let edits = []

	let scanBucket = bucketName => {
		let providerName = bucketName.replace('Match', '').replace('AnimeTo', '')
		return arn.forEach(bucketName, record => {
			record.providerName = providerName

			if(record.edited)
				edits.push(record)
		})
	}

	yield [
		scanBucket('MatchMyAnimeList'),
		scanBucket('MatchHummingBird'),
		scanBucket('MatchAnimePlanet'),
		scanBucket('AnimeToNyaa')
	]

	edits.sort((a, b) => a.edited < b.edited ? 1 : -1)
	
	if(edits.length > maxLogLength)
		edits.length = maxLogLength
	
	let userTasks = {}
	edits.forEach(edit => {
		if(userTasks.hasOwnProperty(edit.editedBy))
			return
		
		userTasks[edit.editedBy] = arn.get('Users', edit.editedBy)
	})

	let users = yield userTasks

	response.render({
		user,
		edits,
		users
	})
}