'use strict'

exports.get = function*(request, response) {
	if(!arn.auth(request, response, 'admin'))
		return

	let user = request.user
	let statusText = yield arn.execute('aql -c \'show sets\'')

	response.render({
		user,
		statusText
	})
}