'use strict'

let Promise = require('bluebird')

let addYieldHandler = Promise.coroutine(function*() {
	try {
		yield []
		yield {}
	} catch(error) {
		Promise.coroutine.addYieldHandler(function(value) {
			if(Array.isArray(value))
				return Promise.all(value)

			if(typeof value === 'object')
				return Promise.props(value)
		})
	}
})

addYieldHandler()