'use strict'

exports.post = (request, response) => {
	let user = request.user
	let provider = request.body.provider
	let providerId = request.body.providerId

	if(!user) {
		response.end('Not logged in!')
		return
	}

	if(user.role !== 'admin' && user.role !== 'editor') {
		response.end('Not authorized!')
		return
	}

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