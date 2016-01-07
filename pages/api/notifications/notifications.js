'use strict'

let arn = require('../../../lib')

let defaultNotifications = [{
	title: 'You have new notifications',
	message: 'Log in on notify.moe to view them',
	icon: '/images/characters/arn-waifu.png',
	tag: 'not-logged-in'
}]

exports.get = (request, response) => {
	let user = request.user

	if(!user) {
		response.json({
			notifications: defaultNotifications
		})
		return
	}

	// TODO: Fetch stored notifications for the user
	// -> Grab them and send all of them
	// -> Delete all of them in the DB
	arn.getAsync('Notifications', user.id).then(record => {
		console.log(`Notifications for ${user.nick}:`, record.notifications)

		response.json({
			notifications: record.notifications
		})
	}).catch(error => {
		response.json({
			notifications: defaultNotifications
		})
	})
}