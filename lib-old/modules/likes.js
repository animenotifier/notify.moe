arn.getLikeController = type => {
	const elementName = type.slice(0, -1).toLowerCase()
	
	return {
		post: function*(request, response) {
			let user = request.user

			if(!user) {
				response.writeHead(HTTP.BAD_REQUEST)
				response.end('Not logged in')
				return
			}
			
			let id = request.params[0]
			
			if(!id) {
				response.writeHead(HTTP.BAD_REQUEST)
				response.end(`No ${elementName} specified`)
				return
			}
			
			let element = yield arn.db.get(type, id)
			
			if(!element.likes)
				element.likes = []
			
			if(element.likes.indexOf(user.id) !== -1) {
				response.end(`Already liked that ${elementName}`)
				return
			}
			
			element.likes.push(user.id)
			
			yield arn.db.set(type, element.id, {
				likes: element.likes
			})
			
			console.log(`${user.nick} liked the ${elementName} '${id}'`)
			response.end('success')
		}
	}
}

arn.getUnlikeController = type => {
	const elementName = type.slice(0, -1).toLowerCase()
	
	return {
		post: function*(request, response) {
			let user = request.user

			if(!user) {
				response.writeHead(HTTP.BAD_REQUEST)
				response.end('Not logged in')
				return
			}
			
			let id = request.params[0]
			
			if(!id) {
				response.writeHead(HTTP.BAD_REQUEST)
				response.end(`No ${elementName} specified`)
				return
			}
			
			let element = yield arn.db.get(type, id)
			
			if(!element.likes)
				element.likes = []
			
			let index = element.likes.indexOf(user.id)
			
			if(index === -1) {
				response.end(`You did not like this ${elementName} yet`)
				return
			}
			
			element.likes.splice(index, 1)
			
			yield arn.db.set(type, element.id, {
				likes: element.likes
			})
			
			console.log(`${user.nick} unliked the ${elementName} '${id}'`)
			response.end('success')
		}
	}
}