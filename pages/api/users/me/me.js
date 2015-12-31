'use strict'

let arn = require('../../../../lib')

exports.post = function(request, response) {
	if(request.body.function !== 'save') {
		response.end()
		return
	}

	let user = request.user

	if(!user) {
		response.end()
		return
	}

	let key = request.body.key
	let value = request.body.value

	/*if(!user.hasOwnProperty(key)) {
		response.end()
		return
	}*/

	if(key === 'nick') {
		let oldNick = user.nick

		arn.getAsync('NickToUser', value)
		.then(record => {
			response.writeHead(409)
			response.end('Username is already taken.')
		})
		.catch(error => {
			user[key] = value

			Promise.all([
				arn.removeAsync('NickToUser', oldNick),
				arn.setAsync('NickToUser', value, { userId: user.id }),
				arn.setUserAsync(user.id, user)
			])
			.then(() => response.end())
			.catch(error => {
				response.writeHead(409)
				response.end(error.message)
			})
		})
		return
	}

	if(key.indexOf('.') !== -1) {
		let parts = key.split('.')
		let obj = user

		for(let i = 0; i < parts.length - 1; i++) {
			let key = parts[i]

			if(!obj.hasOwnProperty(key))
				obj[key] = {}

			obj = obj[key]
		}

		let lastKey = parts[parts.length - 1]
		obj[lastKey] = value
	} else {
		user[key] = value
	}

	arn.setUserAsync(user.id, user)
		.then(() => response.end())
		.catch(error => {
			console.log(error)
			response.writeHead(409)
			response.end(error.message)
		})
}