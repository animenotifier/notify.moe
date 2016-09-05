let RateLimiter = require('limiter').RateLimiter
let fetchLimiter = new RateLimiter(1, 500)

let checkAvatars = coroutine(function*() {
	console.log(chalk.yellow('✖'), 'Updating user avatars...')

	yield arn.listProviders.AniList.authorize()

	yield arn.forEach('Users', user => {
		if(!arn.isActiveUser(user))
			return

		let gravatarURL = gravatar.url(user.email, {s: '50', r: 'x', d: '404'}, true)

		fetchLimiter.removeTokens(1, () => {
			fetch({
				uri: gravatarURL,
				method: 'GET',
				headers: {
					'User-Agent': 'Anime Release Notifier'
				}
			}).then(body => {
				console.log(chalk.green('✔'), `${user.nick}: Global avatar`)

				arn.set('Users', user.id, { avatar: gravatar.url(user.email) })
			}).catch(error => {
				let tasks = []

				if(user.listProviders.AniList && user.listProviders.AniList.userName)
					tasks.push(arn.listProviders.AniList.getUserImage(user.listProviders.AniList.userName))

				if(user.listProviders.HummingBird && user.listProviders.HummingBird.userName)
					tasks.push(arn.listProviders.HummingBird.getUserImage(user.listProviders.HummingBird.userName))

				if(user.listProviders.MyAnimeList && user.listProviders.MyAnimeList.userName)
					tasks.push(arn.listProviders.MyAnimeList.getUserImage(user.listProviders.MyAnimeList.userName))

				Promise.any(tasks)
				.then(imageURL => {
					// Check if the image really exists on the server
					return fetch({
						uri: imageURL,
						method: 'GET',
						headers: {
							'User-Agent': 'Anime Release Notifier'
						}
					}).then(() => imageURL)
				})
				.then(imageURL => {
					console.log(chalk.green('✔'), `${user.nick}: ${imageURL}`)

					arn.set('Users', user.id, { avatar: imageURL })
				})
				.catch(error => {
					console.log(chalk.red('✖'), `${user.nick}: No avatar`)

					arn.set('Users', user.id, { avatar: '' })
				})
			})
		})
	})
})

arn.repeatedly(30 * minutes, checkAvatars)