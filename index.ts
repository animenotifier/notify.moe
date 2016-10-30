import * as arn from 'arn/lib'
declare var global: any

let app = require('aero')()

global.app = app
global.arn = arn
global.HTTP = require('http-status-codes')

app.on('database ready', db => {
	global.db = db
})

// For POST requests
app.use(require('body-parser').json())

// Start the server
app.run()