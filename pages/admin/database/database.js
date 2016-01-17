'use strict'

let Promise = require('bluebird')
let exec = require('child_process').exec

let execute = Promise.promisify((command, callback) => {
    exec(command, function(error, stdout, stderr) {
		callback(error, stdout)
	})
})

exports.get = (request, response) => {
	if(!arn.auth(request, response, 'editor'))
		return

	let user = request.user

	Promise.props({
		statusText: execute('aql -c \'show sets\'')
	}).then(result => {
		response.render({
			user,
			statusText: result.statusText
		})
	})
}