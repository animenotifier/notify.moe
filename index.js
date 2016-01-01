'use strict'

let aero = require('aero')
let fs = require('fs')
let bodyParser = require('body-parser')

// Start the server
aero.run()

// Rewrite URLs
aero.preRoute(function(request, response) {
	if(request.url.startsWith('/+'))
		request.url = '/user/' + request.url.substring(2)
	else if(request.url.startsWith('/_/+'))
		request.url = '/_/user/' + request.url.substring(4)
})

// For POST requests
aero.use(bodyParser.json())

// Load all modules
fs.readdirSync('modules').forEach(mod => require('./modules/' + mod)(aero))