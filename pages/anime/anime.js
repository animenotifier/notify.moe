const genreIcons = {
	'Action': 'bomb',
	'Adventure': 'diamond',
	'Comedy': 'smile-o',
	'Drama': 'heartbeat',
	'Ecchi': 'heart-o',
	'Fantasy' : 'tree',
	'Harem' : 'group',
	'Historical' : 'history',
	'Horror' : 'frown-o',
	'Martial Arts' : 'hand-rock-o',
	'Magic': 'magic',
	'Mecha' : 'reddit-alien',
	'Military' : 'fighter-jet',
	'Mystery' : 'question',
	'Psychological': 'lightbulb-o',
	'Romance': 'heart',
	'Sci-Fi' : 'space-shuttle',
	'School' : 'graduation-cap',
	'Seinen' : 'male',
	'Shounen' : 'male',
	'Shoujo': 'female',
	'Slice of Life' : 'hand-peace-o',
	'Sports': 'soccer-ball-o',
	'Supernatural': 'magic',
	'Super Power': 'flash',
	'Thriller': 'hourglass-end',
	'Vampire': 'eye'
}

exports.get = function*(request, response) {
	let user = request.user
	let animeId = parseInt(request.params[0])
	let category = request.params[1]
	let categoryParameter = request.params[2]

	if(isNaN(animeId)) {
		let popularAnime = yield arn.get('Cache', 'popularAnime')

		return response.render({
			user,
			popularAnime,
			animeToIdCount: arn.animeToIdCount,
			anime: null
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
		
		// For the layout
		request.animePage = animePage

		response.render(Object.assign({
			user,
			videoParameters,
			fixGenre: arn.fixGenre,
			nyaa: arn.animeProviders.Nyaa,
			genreIcons,
			canEdit: user && (user.role === 'admin' || user.role === 'editor')
		}, animePage))
	} catch(error) {
		console.error(error)

		response.render({
			user,
			error: 404
		})
	}
}