'use strict'


let Promise = require('bluebird')

const listLength = 15

exports.get = (request, response) => {
	if(!arn.auth(request, response, 'editor'))
		return

	let user = request.user

	let providers = [
		'MyAnimeList',
		'HummingBird',
		'AnimePlanet'
	]

	let matches = {}
	let tasks = []

	providers.forEach(provider => {
		matches[provider] = []

		tasks.push(arn.scan('Match' + provider, record => {
			if(!record.edited)
				matches[provider].push(record)
		}))
	})

	Promise.all(tasks).then(() => {
		providers.forEach(provider => {
			matches[provider].sort((a, b) => a.similarity > b.similarity ? 1 : -1)

			if(matches[provider].length >= listLength)
				matches[provider].length = listLength
		})

		response.render({
			user,
			matches,
			linkPrefixes: {
				MyAnimeList: 'http://myanimelist.net/anime/',
				HummingBird: 'https://hummingbird.me/anime/',
				AnimePlanet: 'http://anime-planet.com/anime/'
			}
		})
	})
}