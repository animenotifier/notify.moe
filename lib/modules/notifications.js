let Promise = require('bluebird')
let webPush = require('web-push')

webPush.setGCMAPIKey(arn.apiKeys.gcm.serverKey)

arn.sendNotification = (user, notification) => {
	if(!arn.production)
		return Promise.resolve()

	console.log(`Sending notification to ${user.nick}`, notification)

	let notify = Promise.coroutine(function*(record) {
		yield arn.set('Notifications', user.id, record)
		yield Promise.delay(100)

		Object.keys(user.pushEndpoints).forEach(endpoint => {
			webPush.sendNotification(endpoint, {
				TTL: 24 * 60 * 60 * 60
			})
		})
	})

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