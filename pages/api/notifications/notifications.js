'use strict'

exports.get = (request, response) => {
	let user = request.user

	response.json({
		notification: {
			title: user ? (user.nick + ', there is a new episode') : 'Yay a message.',
			message: 'We have received a push message.',
			icon: '/images/characters/arn-waifu.png',
			tag: 'demo'
		}
	})
}