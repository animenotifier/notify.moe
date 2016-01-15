'use strict'

exports.post = (request, response) => {
	if(!arn.auth(request, response, 'editor'))
		return

	let user = request.user
	let provider = request.body.provider
	let providerId = request.body.providerId

	if(!provider || !providerId) {
		response.end('Invalid data!')
		return
	}

	if(provider !== 'AnimePlanet')
		providerId = parseInt(providerId)

	console.log(`${user.nick} confirmed a match for ${provider} ID ${providerId}`)

	arn.set('Match' + provider, providerId, {
		edited: (new Date()).toISOString(),
		editedBy: user.id
	}).finally(() => {
		response.end()
	})
}