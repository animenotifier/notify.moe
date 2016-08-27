let gravatar = require('gravatar')

const emailRegEx = /(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\[(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?|[a-z0-9-]*[a-z0-9]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)\])/

exports.post = function*(request, response) {
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
			response.writeHead(HTTP.BAD_REQUEST)
			response.end('Username must have a length of at least 2 characters')
			return
		}

		return arn.changeNick(user, value)
		.then(() => response.end(user.nick))
		.catch(error => {
			response.writeHead(HTTP.BAD_REQUEST)
			response.end(user.nick)
		})
	} else if((key.startsWith('listProviders.') && key.endsWith('.userName')) || key === 'twitter' || key === 'osu') {
		value = arn.fixListProviderUserName(value)
	} else if(key.startsWith('providers.')) {
		// arn.userListCache.flushAll()
	}
	
	if(key === 'email') {
		if(value.endsWith('googlemail.com'))
			value = value.replace('googlemail.com', 'gmail.com')
		
    	if(!emailRegEx.test(value)) {
			response.writeHead(HTTP.BAD_REQUEST)
			response.end('Invalid email')
			return
		}
		
		try {
			yield arn.get('EmailToUser', value)
			
			response.writeHead(HTTP.BAD_REQUEST)
			response.end('Email already used by another user')
			return
		} catch(error) {
			if(user.email)
				yield arn.remove('EmailToUser', user.email)
			
			yield arn.set('EmailToUser', value, {
				userId: user.id
			})
			
			user.avatar = gravatar.url(value)
		}
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
			console.error(error)
			response.writeHead(HTTP.BAD_REQUEST)
			response.end(error.message)
		})
}