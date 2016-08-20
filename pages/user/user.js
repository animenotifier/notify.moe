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
	let viewUser = null

	try {
		if(!viewUserNick)
			viewUser = user
		else
			viewUser = yield arn.getUserByNick(viewUserNick)
		
		if(!viewUser) {
			response.render({
				user,
				viewUser
			})
			return
		}
		
		viewUser.gravatarURL = gravatar.url(viewUser.email, {s: '320', r: 'x', d: 'mm'}, true)
		
		if(viewUser.role === 'editor')
			viewUser.dataEditCount = (yield arn.get('Cache', 'dataEditCount')).contributions[viewUser.id]
		
		// Open Graph
		request.og = {
			url: app.package.homepage + '/+' + viewUser.nick,
			title: viewUser.nick,
			description: `${viewUser.nick}'s profile`,
			image: viewUser.gravatarURL
		}

		response.render({
			user,
			viewUser,
			embeddedList,
			monthNames
		})
	} catch(error) {
		console.error(error)
		
		response.render({
			user,
			viewUser
		})
	}
}
