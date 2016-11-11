import * as arn from 'arn'
import * as HTTP from 'http-status-codes'

exports.get = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}

	let unfollowUserNick = request.params[0]

	if(!unfollowUserNick) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Username required')
		return
	}

	let unfollowUser = yield arn.getUserByNick(unfollowUserNick)
	let index = user.following.indexOf(unfollowUser.id)

	if(index !== -1) {
		user.following.splice(index, 1)
		yield arn.db.set('Users', user.id, user)
	}

	response.end('success')
}