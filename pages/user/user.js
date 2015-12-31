'use strict'

let gravatar = require('gravatar')
let arn = require('../../lib')

function getProperty(obj, desc) {
	let arr = desc.split('.')

	while(arr.length && obj)
		obj = obj[arr.shift()]

	return obj
}

exports.get = function(request, response) {
	let user = request.user
	let viewUserNick = request.params[0]

	if(!viewUserNick) {
		response.render({ user })
		return
	}

	arn.getUserByNickAsync(viewUserNick)
	.then(viewUser => {
		viewUser.gravatarURL = gravatar.url(viewUser.email, {s: '320', r: 'x', d: 'mm'}, true)
		response.render({
			user,
			viewUser,
			getProperty
		})
	}).catch(error => {
		console.error(error)
		response.render({ user })
		return
	})
}