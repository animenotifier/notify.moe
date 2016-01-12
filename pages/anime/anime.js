'use strict'

let arn = require('../../lib')
let Promise = require('bluebird')

exports.get = function(request, response) {
	let user = request.user
	let animeId = parseInt(request.params[0])

	if(animeId) {
		let providers = {}

		let createScanFunction = function(listProviderName) {
			return match => {
				if(match.id === animeId)
					providers[listProviderName] = match
			}
		}

		let otherProviderTasks = {
			HummingBird: arn.scan('MatchHummingBird', createScanFunction('HummingBird')),
			MyAnimeList: arn.scan('MatchMyAnimeList', createScanFunction('MyAnimeList')),
			AnimePlanet: arn.scan('MatchAnimePlanet', createScanFunction('AnimePlanet'))
		}

		arn.get('Anime', animeId).then(anime => {
			Promise.props(otherProviderTasks).then(result => {
				response.render({
					user,
					anime,
					providers,
					canEdit: user && (user.role === 'admin' || user.role === 'editor')
				})
			})
		}).catch(error => {
			console.log(error.stack)
			response.writeHead(404)
			response.end('Anime not found.')
		})
		return
	}

	response.render({
		user,
		animeToIdJSONString: arn.animeToIdJSONString
	})
}