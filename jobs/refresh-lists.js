// Check every now and then if users have new episodes
let refreshAnimeLists = coroutine(function*() {
	console.log(chalk.yellow('✖'), 'Refreshing anime lists...')
	
	yield arn.listProviders.AniList.authorize()
	
	let users = yield arn.filter('Users', user => arn.isActiveUser(user))
	console.log(`Refreshing anime lists of ${users.length} users`)
	
	for(let user of users) {
		yield Promise.delay(50)
		
		console.log(`Refreshing anime list of ${user.nick}`)
		
		yield arn.getAnimeList(user, true).then(animeList => {
			// ...
		}).catch(error => {
			if(error.name === 'StatusCodeError') {
				console.warn(`Unavailable [${error.statusCode}]: ${error.options.uri}`)
				return
			}
			
			console.error(`Error when automatically updating the anime list of ${user.nick}:`, error, error.stack)
		})
	}
})

arn.repeatedly(arn.animeListCacheTime / 1000, refreshAnimeLists)