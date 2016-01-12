'use strict'

let arn = require('../../lib')
let Promise = require('bluebird')
let exec = require('child_process').exec

const listLength = 15

let execute = Promise.promisify((command, callback) => {
    exec(command, function(error, stdout, stderr) {
		callback(error, stdout)
	})
})

exports.get = (request, response) => {
	let user = request.user
	let malMatches = []
	let hbMatches = []
	let apMatches = []

	Promise.props({
		statusText: execute('sugoi stats'),
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
		let status = result.statusText.split('\n').map(line => line.split(':').map(value => value.trim()))

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
			status,
			malMatches,
			hbMatches,
			apMatches
		})
	})
}