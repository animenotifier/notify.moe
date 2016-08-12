let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
    arn.forEach('Users', function(user) {
        user.devices = {}
		arn.set('Users', user.id, user)
		console.log(user.nick)
    }).then(function() {
        console.log('Finished updating all users')
    })
})