import * as arn from 'arn'

global.arn = arn
global.app = require('aero')()

app.on('database ready', db => {
	global.db = db
})

// For POST requests
app.use(require('body-parser').json())

// Start the server
app.run()