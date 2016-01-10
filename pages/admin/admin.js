'use strict'

let arn = require('../../lib')
let Promise = require('bluebird')
let exec = require('child_process').exec

let execute = Promise.promisify((command, callback) => {
    exec(command, function(error, stdout, stderr) {
		callback(error, stdout)
	})
})

exports.get = (request, response) => {
	let user = request.user

	execute('sugoi stats').then(statusText => {
		let status = statusText.split('\n').map(line => line.split(':').map(value => value.trim()))
		response.render({
			user,
			status
		})
	})
}