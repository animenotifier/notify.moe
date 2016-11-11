import * as arn from 'arn'
import * as HTTP from 'http-status-codes'

exports.post = async function(request, response) {
	let user = request.user

	if(!user) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('Not logged in')
		return
	}

	let messageId = request.params[0]

	if(!messageId) {
		response.writeHead(HTTP.BAD_REQUEST)
		response.end('No message specified')
		return
	}

	// TODO: Check message author ID

	await arn.db.remove('Messages', messageId)

	console.log(`${user.nick} deleted the message '${messageId}'`)
	response.end('success')
}