'use strict'


let Promise = require('bluebird')
let exec = require('child_process').exec

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
		statusText: execute('sugoi stats')
	}).then(result => {
		let status = result.statusText.split('\n').map(line => line.split(':').map(value => value.trim()))

		response.render({
			user,
			status
		})
	})
}