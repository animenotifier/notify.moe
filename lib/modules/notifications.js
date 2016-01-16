'use strict'

let gcm = require('node-gcm')
let sender = new gcm.Sender(arn.apiKeys.gcm.serverKey)

arn.sendNotification = (user, notification) => {
	let message = new gcm.Message({
	    notification, // Note that this is actually ignored by GCM
		delayWhileIdle: false
	})
	message.addData('test', 'test')

	console.log(`Sending notification to ${user.nick}`)

	let notify = record => {
		return arn.set('Notifications', user.id, record).then(() => {
			// Now let the user know that we have a notification saved for him
			return sender.send(message, {
				registrationTokens: Object.keys(user.devices)
			}, 10, function(err, response) {
				if(err)
					console.error(err)
				else
					console.log(response)
			})
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