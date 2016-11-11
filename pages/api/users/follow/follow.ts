import * as arn from 'arn'
import * as HTTP from 'http-status-codes'

exports.get = async function(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}

	let followUserNick = request.params[0]

	if(!followUserNick) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Username required')
		return
	}

	let followUser = await arn.getUserByNick(followUserNick)

	if(user.id !== followUser.id && user.following.indexOf(followUser.id) === -1) {
		user.following.push(followUser.id)
		await arn.db.set('Users', user.id, user)
	}

	response.end('success')
}