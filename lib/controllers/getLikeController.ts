import * as arn from 'arn'

const HTTP = require('http-status-codes')

export function getLikeController(type: string) {
	const elementName = type.slice(0, -1).toLowerCase()

	return {
		post: async (request, response) => {
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

			if(element.likes.indexOf(user.id) !== -1) {
				response.end(`Already liked that ${elementName}`)
				return
			}

			element.likes.push(user.id)

			await arn.db.set(type, element.id, {
				likes: element.likes
			})

			console.log(`${user.nick} liked the ${elementName} '${id}'`)
			response.end('success')
		}
	}
}