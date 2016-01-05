'use strict'

let gcm = require('node-gcm')
let apiKeys = require('../security/api-keys.json')
let sender = new gcm.Sender(apiKeys.gcm.serverKey)

exports.send = (user, notification) => {
	let message = new gcm.Message({
	    notification // Note that this is actually ignored by GCM
	})

	console.log(`Sending notification to ${user.nick}`)

	// TODO: Save notification in database
	// ...

	// Now let the user know that we have a notification saved for him
	sender.sendNoRetry(message, {
		registrationTokens: Object.keys(user.devices)
	}, function(err, response) {
		if(err)
			console.error(err)
		else
			console.log(response)
	})
}