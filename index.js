'use strict'

let aero = require('aero')
let fs = require('fs')
let bodyParser = require('body-parser')

// Start the server
aero.run()

// For POST requests
aero.use(bodyParser.json())

// Load all modules
fs.readdirSync('modules').forEach(mod => require('./modules/' + mod)(aero))