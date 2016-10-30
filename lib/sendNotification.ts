import * as arn from '.'
import * as Promise from 'bluebird'
import { User } from './interfaces/User'
import { Notification } from './interfaces/Notification'
const webPush = require('web-push')

const vapid = webPush.generateVAPIDKeys()

webPush.setGCMAPIKey(arn.api.gcm.serverKey)
webPush.setVapidDetails('mailto:push@notify.moe', vapid.publicKey, vapid.privateKey)

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