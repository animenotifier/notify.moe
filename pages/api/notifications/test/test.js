'use strict'



exports.get = (request, response) => {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}

	arn.sendNotification(user, {
		title: 'Anime Title [123]',
		icon: 'https://notify.moe/images/characters/arn-waifu.png',
		message: 'New episode available',
		tag: 'test'
	})
}