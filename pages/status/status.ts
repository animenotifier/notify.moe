import * as arn from 'arn'

exports.get = async function(request, response) {
	let status = await arn.db.get('Cache', 'status')

	response.render({
		user: request.user,
		status
	})
}