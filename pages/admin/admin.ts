import * as arn from 'arn'

exports.get = async function(request, response) {
	if(!arn.auth(request, response, 'editor'))
		return

	let user = request.user
	let statusText = await arn.execute('sugoi stats')
	let status = statusText.trim().split('\n').map(line => line.split(':').map(value => value.trim()))

	response.render({
		user,
		status
	})
}