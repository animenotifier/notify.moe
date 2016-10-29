exports.get = function*(request, response) {
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
	
	let followUser = yield arn.getUserByNick(followUserNick)
	
	if(user.id !== followUser.id && user.following.indexOf(followUser.id) === -1) {
		user.following.push(followUser.id)
		yield arn.db.set('Users', user.id, user)
	}
	
	response.end('success')
}