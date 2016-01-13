'use strict'

let arn = require('../../../lib')
let Promise = require('bluebird')

const listLength = 15

exports.get = (request, response) => {
	let user = request.user

	if(!user) {
		response.render({})
		return
	}

	let edits = []

	let scanBucket = bucketName => {
		let providerName = bucketName.replace('Match', '')
		return arn.scan(bucketName, record => {
			record.providerName = providerName

			if(record.edited)
				edits.push(record)
		})
	}

	Promise.all([
		scanBucket('MatchMyAnimeList'),
		scanBucket('MatchHummingBird'),
		scanBucket('MatchAnimePlanet')
	]).then(() => {
		edits.sort((a, b) => a.edited < b.edited ? 1 : -1)

		let getUsers = edits.map(edit => {
			return arn.get('Users', edit.editedBy).then(editor => edit.editedBy = editor)
		})

		Promise.all(getUsers).then(() => {
			response.render({
				user,
				edits
			})
		})
	})
}