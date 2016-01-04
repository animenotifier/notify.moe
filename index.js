'use strict'

let aero = require('aero')
let arn = require('./lib')
let fs = require('fs')
let bodyParser = require('body-parser')
let request = require('request-promise')

// Start the server
aero.run()

// Rewrite URLs
aero.preRoute(function(request, response) {
	if(request.headers.host.indexOf('animereleasenotifier.com') !== -1) {
        response.redirect('https://notify.moe' + request.url)
        return true
    }

	if(request.url.startsWith('/+'))
		request.url = '/user/' + request.url.substring(2)
	else if(request.url.startsWith('/_/+'))
		request.url = '/_/user/' + request.url.substring(4)
})

// For POST requests
aero.use(bodyParser.json())

// Send slack messages
arn.on('new user', function(user) {
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
		console.error('Error sending slack message:', error)
	})
})

arn.on('new forum reply', function(link, userName) {
	let webhook = 'https://hooks.slack.com/services/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'

	request.post({
		url: webhook,
		body: JSON.stringify({
			text: `<${link}|${userName}>`
		})
	}).then(body => {
		console.log(`Sent slack message about a new forum reply from ${userName}`)
	}).catch(error => {
		console.error('Error sending slack message:', error)
	})
})

// Load all modules
fs.readdirSync('modules').forEach(mod => require('./modules/' + mod)(aero))