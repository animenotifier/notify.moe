'use strict'

exports.get = function*(request, response) {
	let user = request.user
	let animeId = parseInt(request.params[0])
	let category = request.params[1]
	let categoryParameter = request.params[2]

	console.log(request.params)

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
		let videoParameters = ''

		if(category === 'video') {
			videoParameters += '&autoplay=1'

			if(categoryParameter && categoryParameter.indexOf('s') !== -1) {
				videoParameters += '&start=' + categoryParameter.replace('s', '')
			}
		}

		response.render(Object.assign({
			user,
			videoParameters,
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