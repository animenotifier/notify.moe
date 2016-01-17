'use strict'

let natural = require('natural')

let editListProviderId = (request, response, user, animeId) => {
	let bucket = 'Match' + request.body.key
	let providerId = request.body.value.trim()
	let oldProviderId = request.body.old

	if(providerId === oldProviderId) {
		response.end('Can not change to same ID!')
		return
	}

	if(bucket !== 'MatchAnimePlanet') {
		providerId = parseInt(providerId)
		oldProviderId = parseInt(oldProviderId)
	}

	arn.get('Anime', animeId).then(anime => {
		if((!providerId || providerId === NaN) && (oldProviderId && oldProviderId !== NaN)) {
			console.log(`${user.nick} deleted ${request.body.key} ID of '${anime.title.romaji}' (https://notify.moe/anime/${anime.id}): ${oldProviderId} => DELETED`)
			return arn.remove(bucket, oldProviderId).finally(() => {
				response.end()
			})
		}

		let removeOldAndSave = match => {
			let verb = 'edited'
			if(!oldProviderId || oldProviderId === NaN)
				verb = 'added'
			console.log(`${user.nick} ${verb} ${request.body.key} ID of '${anime.title.romaji}' (https://notify.moe/anime/${anime.id}):`, verb === 'added' ? providerId : `${oldProviderId} => ${providerId}`)

			if(oldProviderId !== NaN)
				return arn.remove(bucket, oldProviderId).catch(error => null).finally(() => arn.set(bucket, providerId, match))
			else
				return arn.set(bucket, providerId, match)
		}

		return arn.get(bucket, providerId).then(match => {
			match.id = anime.id
			match.title = anime.title.romaji
			match.similarity = natural.JaroWinklerDistance(match.providerTitle, match.title)
			match.edited = (new Date()).toISOString()
			match.editedBy = user.id
			return match
		}).then(match => {
			return removeOldAndSave(match)
		}).catch(error => {
			if(error && error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND') {
				return removeOldAndSave({
					id: anime.id,
					title: anime.title.romaji,
					similarity: null,
					providerId,
					providerTitle: '',
					edited: (new Date()).toISOString(),
					editedBy: user.id
				})
			} else {
				console.error(error.stack)
			}
		}).finally(() => {
			response.end()
		})
	}).catch(error => {
		if(error && error.message === 'AEROSPIKE_ERR_RECORD_NOT_FOUND') {
			console.error('Anime not found: ' + animeId)
		} else {
			console.error(error.stack)
		}
		response.end()
	})
}

exports.post = (request, response) => {
	if(!arn.auth(request, response, 'editor'))
		return

	let user = request.user
	let animeId = parseInt(request.params[0])

	if(!animeId) {
		response.end('Invalid ID!')
		return
	}

	let provider = request.body.key

	if(provider === 'Nyaa') {
		let old = request.body.old
		let title = request.body.value

		if(title) {
			console.log(`${user.nick} set Nyaa title of https://notify.moe/anime/${animeId} from '${old}' to '${title}'`)

			arn.set('AnimeToNyaa', animeId, {
				id: animeId,
				title,
				edited: (new Date()).toISOString(),
				editedBy: user.id
			}).then(() => response.end())
		} else {
			console.log(`${user.nick} deleted Nyaa title of https://notify.moe/anime/${animeId} which was '${old}'`)

			arn.remove('AnimeToNyaa', animeId).then(() => response.end())
		}
	} else {
		editListProviderId(request, response, user, animeId)
	}
}