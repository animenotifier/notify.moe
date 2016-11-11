const hosts = [
	'https://anilist.co',
	'https://hummingbird.me',
	'http://myanimelist.net',
	'http://www.anime-planet.com',
	'http://www.crunchyroll.com',
	'https://www.nyaa.se'
]

let updateStatus = coroutine(function*() {
	let status = []
	
	console.log(chalk.cyan('↻'), 'Updating status...')
	
	for(let host of hosts) {
		yield Promise.delay(100)
		
		yield fetch({
			uri: host,
			method: 'GET',
			headers: {
				'User-Agent': 'Anime Release Notifier'
			}
		})
		.then(body => {
			console.log(chalk.green('✔'), host, 'alive')
			
			status.push({
				context: host,
				error: ''
			})
		})
		.catch(e => {
			console.log(chalk.red('✖'), host, 'dead')
			
			status.push({
				context: host,
				error: e.toString()
			})
		})
	}
	
	arn.db.set('Cache', 'status', status)
})

arn.repeatedly(5 * minutes, updateStatus)