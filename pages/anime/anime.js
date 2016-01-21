'use strict'

exports.get = function*(request, response) {
	let user = request.user
	let animeId = parseInt(request.params[0])

	if(!animeId) {
		let popularAnime = yield arn.get('Cache', 'popularAnime')

		return response.render({
			user,
			popularAnime,
			animeToIdCount: arn.animeToIdCount
		})
	}

	try {
		let animePage = yield arn.get('AnimePages', animeId)

		response.render(Object.assign({
			user,
			fixGenre: arn.fixGenre,
			nyaa: arn.animeProviders.Nyaa,
			canEdit: user && (user.role === 'admin' || user.role === 'editor')
		}, animePage))
	} catch(error) {
		console.error(error, error.stack)

		response.render({
			user,
			error: 404
		})
	}
}