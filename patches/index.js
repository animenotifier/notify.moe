'use strict'

let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
	arn.getUserByNickAsync('Aky').then(user => {
		user.role = 'admin'

		arn.setUserAsync(user.id, user).then(() => {
			console.log('Finished updating ' + user.nick)
		})
	})
})