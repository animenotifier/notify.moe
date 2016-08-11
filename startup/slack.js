let request = require('request-promise')

// Send slack messages
arn.on('new user', user => {
	// Ignore my own attempts on empty databases
	if(user.email === 'e.urbach@gmail.com')
		return

	let host = 'https://notify.moe'
	let webhook = 'https://hooks.slack.com/services/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'

	request.post({
		url: webhook,
		body: JSON.stringify({
			text: `<${host}/users|${user.firstName} ${user.lastName} (${user.email})>`
		})
	}).then(body => {
		console.log(`Sent slack message about the new user registration: ${user.email}`)
	}).catch(error => {
		console.error('Error sending slack message:', error, error.stack)
	})
})