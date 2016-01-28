'use strict'

let chalk = require('chalk')
let request = require('request-promise')
let RateLimiter = require('limiter').RateLimiter

let osuAPILimiter = new RateLimiter(1, 100)

let updateOsuDetails = function() {
	console.log(chalk.yellow('✖'), 'Updating osu ranks...')

	arn.forEach('Users', user => {
		if(!user.osu)
			return

		let apiURL = `https://osu.ppy.sh/api/get_user?k=${arn.apiKeys.osu.clientSecret}&u=${user.osu}`

		osuAPILimiter.removeTokens(1, () => {
			request({
				uri: apiURL,
				method: 'GET',
				headers: {
					'User-Agent': 'Anime Release Notifier',
					'Accept': 'application/json'
				}
			}).then(body => {
				let osu = JSON.parse(body)[0]

				arn.set('Users', user.id, {
					osuDetails: {
						nick: osu.username,
						pp: parseFloat(osu.pp_raw),
						level: parseFloat(osu.level),
						accuracy: parseFloat(osu.accuracy),
						playCount: parseInt(osu.playcount)
					}
				})
			}).catch(error => {
				console.error(error, error.stack)
			})
		})
	})
}

arn.repeatedly(60 * 60, () => {
	updateOsuDetails()
})