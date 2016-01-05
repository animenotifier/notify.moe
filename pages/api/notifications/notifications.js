'use strict'

exports.get = (request, response) => {
	response.json({
		notification: {
			title: 'Yay a message.',
			message: 'We have received a push message.',
			icon: '/images/characters/arn-waifu.png',
			tag: 'demo'
		}
	})
}