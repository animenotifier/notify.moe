global.arn = require('./lib')
global.app = require('aero')()
global.HTTP = require('http-status-codes')

// For POST requests
app.use(require('body-parser').json())

// Start the server
app.run()