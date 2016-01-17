'use strict'

let Promise = require('bluebird')

exports.get = (request, response) => {
	if(!arn.auth(request, response, 'editor'))
		return

	let user = request.user

	let providers = [
		'Nyaa'
	]

	let matches = providers.reduce((obj, provider) => {
		obj[provider] = []
		return obj
	}, {})

	let tasks = providers.map(provider => arn.filter('AnimeTo' + provider, record => !record.edited && record.episodes === 0).then(uneditedMatches => matches[provider] = uneditedMatches))

	Promise.all(tasks).then(() => {
		/*providers.forEach(provider => {
			matches[provider].sort((a, b) => a.similarity > b.similarity ? 1 : -1)

			if(matches[provider].length >= listLength)
				matches[provider].length = listLength
		})*/

		response.render({
			user,
			matches
		})
	})
}