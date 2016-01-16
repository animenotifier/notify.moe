'use strict'

let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
	arn.forEach('Users', function(user) {
		// ...
	}).then(function() {
		console.log('Finished updating all users')
	})
})