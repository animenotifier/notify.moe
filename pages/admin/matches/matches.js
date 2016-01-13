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

	let malMatches = []
	let hbMatches = []
	let apMatches = []

	Promise.props({
		scanMAL: arn.scan('MatchMyAnimeList', record => {
			if(!record.edited)
				malMatches.push(record)
		}),
		scanHB: arn.scan('MatchHummingBird', record => {
			if(!record.edited)
				hbMatches.push(record)
		}),
		scanAP: arn.scan('MatchAnimePlanet', record => {
			if(!record.edited)
				apMatches.push(record)
		})
	}).then(result => {
		malMatches.sort((a, b) => a.similarity > b.similarity ? 1 : -1)

		if(malMatches.length >= listLength)
			malMatches.length = listLength

		hbMatches.sort((a, b) => a.similarity > b.similarity ? 1 : -1)

		if(hbMatches.length >= listLength)
			hbMatches.length = listLength

		apMatches.sort((a, b) => a.similarity > b.similarity ? 1 : -1)

		if(apMatches.length >= listLength)
			apMatches.length = listLength

		response.render({
			user,
			malMatches,
			hbMatches,
			apMatches
		})
	})
}