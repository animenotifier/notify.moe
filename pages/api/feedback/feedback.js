'use strict'

let $ = require('request-promise')
let webhook = 'https://hooks.slack.com/services/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'

exports.post = (request, response) => {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}

	let feedback = request.body.text

	$.post({
		url: webhook,
		body: JSON.stringify({
			text: `<https://notify.moe/+${user.nick}|${user.nick}>\n>>> ${feedback}`
		})
	}).then(body => {
		response.end('OK')
	}).catch(error => {
		response.end('Error: ' + error)
	})
}