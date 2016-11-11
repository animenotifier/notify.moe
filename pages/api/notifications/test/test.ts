import * as arn from 'arn'
import * as HTTP from 'http-status-codes'

exports.get = (request, response) => {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}

	arn.sendNotification(user, {
		title: 'Anime Title [123]',
		icon: 'https://notify.moe/images/characters/arn-waifu.png',
		body: 'New episode available'
	})
}