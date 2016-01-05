'use strict'

let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
    arn.scan('Users', function(user) {
        user.devices = {}
		arn.setUser(user.id, user)
		console.log(user.nick)
    }, function() {
        console.log('Finished updating all users')
    })
})