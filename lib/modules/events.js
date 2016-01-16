'use strict'

let EventEmitter = require('events').EventEmitter

arn.events = new EventEmitter()

arn.on = function(eventName, func) {
	arn.events.on(eventName, func)
}