let Promise = require('bluebird')

const maxLogLength = 100

exports.get = function*(request, response) {
	if(!arn.auth(request, response, 'editor'))
		return

	let user = request.user
	let edits = []

	let scanBucket = bucketName => {
		let providerName = bucketName.replace('Match', '').replace('AnimeTo', '')
		return arn.db.forEach(bucketName, record => {
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
	
	let userTasks = {}
	let userEdits = {}
	edits.forEach(edit => {
		if(userEdits.hasOwnProperty(edit.editedBy))
			userEdits[edit.editedBy] += 1
		else
			userEdits[edit.editedBy] = 1
		
		if(userTasks.hasOwnProperty(edit.editedBy))
			return
		
		userTasks[edit.editedBy] = arn.db.get('Users', edit.editedBy)
	})

	let users = yield userTasks
	
	// Save editor contribution
	arn.db.set('Cache', 'dataEditCount', {
		contributions: userEdits
	})
	
	// Don't optimize it by putting it before the loop.
	// It would mess up editor edit counts.
	edits.sort((a, b) => a.edited < b.edited ? 1 : -1)
	
	if(edits.length > maxLogLength)
		edits.length = maxLogLength

	response.render({
		user,
		edits,
		users
	})
}