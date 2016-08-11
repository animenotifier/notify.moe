let RateLimiter = require('limiter').RateLimiter
let fetchLimiter = new RateLimiter(1, 500)

let checkAvatars = function() {
	console.log(chalk.yellow('✖'), 'Updating user avatars...')

	arn.forEach('Users', user => {
		let gravatarURL = gravatar.url(user.email, {s: '50', r: 'x', d: '404'}, true)
		
		fetchLimiter.removeTokens(1, () => {
			fetch({
				uri: gravatarURL,
				method: 'GET',
				headers: {
					'User-Agent': 'Anime Release Notifier'
				}
			}).then(body => {
				arn.set('Users', user.id, {
					avatar: gravatar.url(user.email)
				})
			}).catch(error => {
				arn.set('Users', user.id, {
					avatar: ''
				})
			})
		})
	})
}

arn.repeatedly(30 * minutes, checkAvatars)