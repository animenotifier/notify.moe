'use strict'

let gcm = require('node-gcm')

// Use different endpoints for each browser
let chrome = new gcm.Sender(arn.apiKeys.gcm.serverKey)
let firefox = new gcm.Sender(arn.apiKeys.gcm.serverKey, {
	uri: 'https://updates.push.services.mozilla.com/push/'
})

arn.sendNotification = (user, notification) => {
	if(!arn.production)
		return Promise.resolve()

	let message = new gcm.Message({
	    notification, // Note that this is actually ignored by GCM
		priority: 'high',
		delayWhileIdle: false
	})

	console.log(`Sending notification to ${user.nick}`)

	let notify = record => {
		return arn.set('Notifications', user.id, record).then(() => {
			// Chrome
			chrome.send(message, {
				registrationTokens: Object.keys(user.devices)
			}, 5, function(err, response) {
				if(err)
					console.error(err)
				else
					console.log(response)
			})

			// Firefox
			firefox.send(message, {
				registrationTokens: Object.keys(user.devices)
			}, 5, function(err, response) {
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