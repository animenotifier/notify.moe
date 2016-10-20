declare var global: any
declare var app: any

global.app = require('aero')()
global.arn = require('./lib')
global.HTTP = require('http-status-codes')

app.on('database ready', db => {
	global.db = db
})

// For POST requests
app.use(require('body-parser').json())

// Start the server
app.run()