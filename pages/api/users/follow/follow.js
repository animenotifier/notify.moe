exports.post = function*(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(409)
		response.end('Not logged in')
		return
	}

	let followUserNick = request.params[0]

	if(!followUserNick) {
		response.writeHead(409)
		response.end('Username required')
		return
	}
	
	let followUser = yield arn.getUserByNick(followUserNick)
	
	if(user.following.indexOf(followUser.id) === -1)
		user.following.push(followUser.id)
	
	yield arn.set('Users', user.id, user)
	
	response.end('success')
}