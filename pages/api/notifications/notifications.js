'use strict'

exports.get = (request, response) => {
	let user = request.user

	if(!user) {
		response.json({
			notification: {
				title: 'You have new notifications',
				message: 'Log in on notify.moe to view them',
				icon: '/images/characters/arn-waifu.png',
				tag: 'not-logged-in'
			}
		})
		return
	}

	// TODO: Fetch stored notifications for the user
	// -> Grab them and send all of them
	// -> Delete all of them in the DB
	response.json({
		notification: {
			title: 'New episode of [XXX]',
			message: user.nick + ', there is a new episode you can watch now.',
			icon: '/images/characters/arn-waifu.png',
			tag: 'new-episode'
		}
	})
}