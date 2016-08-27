global.arn = require('./lib')
global.app = require('aero')()
global.HTTP = require('http-status-codes')

// For POST requests
app.use(require('body-parser').json())

// Run all startup modules
require('fs').readdirSync('startup').forEach(mod => require('./startup/' + mod))

// Start the server
app.run()