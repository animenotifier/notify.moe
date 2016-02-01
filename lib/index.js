'use strict'

let fs = require('fs')
let path = require('path')
let EventEmitter = require('events').EventEmitter

global.arn = {
	apiKeys: require('../security/api-keys.json'),
	events: new EventEmitter(),
	maintenance: false,
	production: process.env.NODE_ENV === 'production'
}

arn.cacheAnimeLists = arn.production

arn.on = function(eventName, func) {
	arn.events.on(eventName, func)
}

// Load every module inside the lib/modules directory
let modules = fs.readdirSync('./lib/modules')
modules.forEach(file => require('./' + path.join('modules', file.replace('.js', ''))))

module.exports = global.arn