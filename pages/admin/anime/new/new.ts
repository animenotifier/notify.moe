import * as arn from 'arn'

exports.get = (request, response) => {
	if(!arn.auth(request, response, 'editor'))
		return

	response.render({
		user: request.user
	})
}