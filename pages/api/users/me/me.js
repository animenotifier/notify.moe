'use strict'



exports.post = function(request, response) {
	if(request.body.function !== 'save') {
		response.end('Invalid function!')
		return
	}

	let user = request.user

	if(!user) {
		response.end('Not logged in!')
		return
	}

	let key = request.body.key
	let value = request.body.value

	/*if(!user.hasOwnProperty(key)) {
		response.end()
		return
	}*/

	if(key === 'nick') {
		value = arn.fixNick(value)

		if(!value || value.length < 2) {
			response.writeHead(409)
			response.end('Username must have a length of at least 2 characters')
			return
		}

		arn.changeNick(user, value)
		.then(() => response.end(user.nick))
		.catch(error => {
			response.writeHead(409)
			response.end(user.nick)
		})
		return
	} else if(key.startsWith('listProviders.') && key.endsWith('.userName')) {
		value = arn.fixListProviderUserName(value)
	} else if(key.startsWith('providers.')) {
		arn.userListCache.flushAll()
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

	arn.set('Users', user.id, user)
		.then(() => response.end())
		.catch(error => {
			console.error(error.stack)
			response.writeHead(409)
			response.end(error.message)
		})
}