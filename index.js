global.arn = require('./lib')
global.app = require('aero')()
global.HTTP = require('http-status-codes')

if(app.production)
	arn.maintenance = true

app.on('database ready', db => {
	global.db = db
})

// For POST requests
app.use(require('body-parser').json())

let chalk = require('chalk')
app.use((request, response, next) => {
	console.log(chalk.cyan(request.url))
	console.log(request.params)
	console.log(request.query)
	next()
})

// Start the server
app.run()