'use strict'

let aero = require('aero')
let fs = require('fs')

// Start the server
aero.run()

// Load all modules
fs.readdirSync('modules').forEach(mod => require('./modules/' + mod)(aero))