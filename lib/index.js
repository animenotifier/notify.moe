'use strict'

let fs = require('fs')
let aero = require('aero')
let Promise = require('bluebird')
let request = require('request-promise')
let path = require('path')

global.arn = {
	apiKeys: require('../security/api-keys.json'),
	maintenance: false,
	cacheAnimeLists: false
}

// Load every module inside the lib/modules directory
let modules = fs.readdirSync('./lib/modules')
modules.forEach(file => require('./' + path.join('modules', file.replace('.js', ''))))

module.exports = global.arn