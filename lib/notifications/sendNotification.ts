import * as arn from 'arn'
import * as Promise from 'bluebird'
import { User } from 'arn/interfaces/User'
import { Notification } from 'arn/interfaces/Notification'

const webPush = require('web-push')
const vapid = arn.api.vapid

webPush.setGCMAPIKey(arn.api.gcm.serverKey)
webPush.setVapidDetails(vapid.subject, vapid.publicKey, vapid.privateKey)

export function sendNotification(user: User, notification: Notification) {
	if(!arn.production)
		return Promise.resolve()

	console.log(`Sending notification to ${user.nick}`, notification)

	let notify = Promise.coroutine(function*(record) {
		yield arn.db.set('Notifications', user.id, record)
		yield Promise.delay(100)

		Object.keys(user.pushEndpoints).forEach(endpoint => {
			const subscription = {
				endpoint,
				keys: user.pushEndpoints[endpoint].keys
			}

			if(!subscription.endpoint || !subscription.keys)
				return

			webPush.sendNotification(subscription, JSON.stringify(notification), { TTL: 24 * 60 * 60 * 60 })
			.catch(console.error)
		})
	})

	return arn.db.get('Notifications', user.id).then(record => {
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