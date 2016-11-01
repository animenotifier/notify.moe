import * as arn from 'arn'

const HTTP = require('http-status-codes')

export function getUnlikeController(type: string) {
	const elementName = type.slice(0, -1).toLowerCase()

	return {
		post: async function(request, response) {
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

			let element = await arn.db.get(type, id)

			if(!element.likes)
				element.likes = []

			let index = element.likes.indexOf(user.id)

			if(index === -1) {
				response.end(`You did not like this ${elementName} yet`)
				return
			}

			element.likes.splice(index, 1)

			await arn.db.set(type, element.id, {
				likes: element.likes
			})

			console.log(`${user.nick} unliked the ${elementName} '${id}'`)
			response.end('success')
		}
	}
}