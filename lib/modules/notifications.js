'use strict'

let webPush = require('web-push')

webPush.setGCMAPIKey(arn.apiKeys.gcm.serverKey)

arn.sendNotification = (user, notification) => {
	if(!arn.production)
		return Promise.resolve()

	console.log(`Sending notification to ${user.nick}`)

	let notify = record => {
		return arn.set('Notifications', user.id, record).then(() => {
			Object.keys(user.pushEndpoints).forEach(webPush.sendNotification)
		})
	}

	return arn.get('Notifications', user.id).then(record => {
		// Add to existing, queued-up notifications
		record.notifications.push(notification)
		return notify(record)
	}).catch(error => {
		// Create first queued-up notification
		return notify({
			notifications: [notification]
		})
	})
}