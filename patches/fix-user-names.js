'use strict'

let aero = require('aero')
let database = require('../modules/database')
let arn = require('../lib')

database(aero, function(error) {
	let tasks = []

    arn.scan('Users', function(user) {
		let listProviderName = user.providers.list
        if(user.listProviders[listProviderName] && user.listProviders[listProviderName].userName) {
            let userName = user.listProviders[listProviderName].userName

			user.listProviders[listProviderName].userName = arn.fixListProviderUserName(userName)
			console.log(user.listProviders[listProviderName].userName)

			tasks.push(arn.setUser(user.id, user))
        }
    }).then(function() {
		Promise.all(tasks).then(() => console.log(`Finished updating ${tasks.length} users`))
    })
})