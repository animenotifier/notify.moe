let gravatar = require('gravatar')

const monthNames = [
	'January',
	'February',
	'March',
	'April',
	'May',
	'June',
	'July',
	'August',
	'September',
	'October',
	'November',
	'December'
]

exports.get = function*(request, response) {
	let user = request.user
	let viewUserNick = request.params[0]
	let embeddedList = request.params[1] === 'watching'

	if(!viewUserNick) {
		let viewUser = user

		if(viewUser) {
			viewUser.gravatarURL = gravatar.url(viewUser.email, {s: '320', r: 'x', d: 'mm'}, true)
			
			if(viewUser.role === 'editor')
				viewUser.dataEditCount = (yield arn.get('Cache', 'dataEditCount')).contributions[viewUser.id]
		}

		response.render({
			user,
			viewUser,
			embeddedList,
			monthNames
		})
		return
	}
	
	// Very old Android app requests
	if(viewUserNick.indexOf('&animeProvider=') !== -1) {
		response.writeHead(409)
		response.end()
		return
	}

	try {
		let viewUser = yield arn.getUserByNick(viewUserNick)
		viewUser.gravatarURL = gravatar.url(viewUser.email, {s: '320', r: 'x', d: 'mm'}, true)
		
		if(viewUser.role === 'editor')
			viewUser.dataEditCount = (yield arn.get('Cache', 'dataEditCount')).contributions[viewUser.id]

		response.render({
			user,
			viewUser,
			embeddedList,
			monthNames
		})
	} catch(error) {
		console.error(error, error.stack)
		response.render({
			user
		})
	}
}